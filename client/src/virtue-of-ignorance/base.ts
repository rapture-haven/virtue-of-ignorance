import EventEmitter from 'eventemitter3'
import { EVENT, WebSocketEvents } from './events'

import {
    WebSocketMessages,
    WebSocketPayloads,
    SignalProvidePayload,
    SignalCandidatePayload,
    SignalOfferPayload,
    SignalAnswerMessage,
} from './messages'

export interface BaseEvents {
    info: (...message: any[]) => void
    warn: (...message: any[]) => void
    debug: (...message: any[]) => void
    error: (error: Error) => void
}

export abstract class BaseClient extends EventEmitter<BaseEvents> {
    protected _ws?: WebSocket
    protected _peer?: RTCPeerConnection
    protected _channel?: RTCDataChannel
    protected _timeout?: number
    protected _displayname?: string
    protected _state: RTCIceConnectionState = 'disconnected'
    protected _id = ''
    protected _candidates: RTCIceCandidate[] = []

    get id() {
        return this._id
    }

    get supported() {
        return typeof RTCPeerConnection !== 'undefined' && typeof RTCPeerConnection.prototype.addTransceiver !== 'undefined'
    }

    get socketOpen() {
        return typeof this._ws !== 'undefined' && this._ws.readyState === WebSocket.OPEN
    }

    get peerConnected() {
        return typeof this._peer !== 'undefined' && ['connected', 'checking', 'completed'].includes(this._state)
    }

    get connected() {
        return this.peerConnected && this.socketOpen
    }

    public connect(url: string, password: string, displayname: string) {
        // TODO: check socket open and support webrtc

        this._displayname = displayname;
        this[EVENT.CONNECTING]();

        // TODO: connect
        throw new Error('Method not implemented.')
    }

    protected disconnect() {
        this._state = 'disconnected'
        this._displayname = undefined
        this._id = ''
        throw new Error('Method not implemented.')
    }

    public sendData(event: 'wheel' | 'mousemove', data: { x: number; y: number }): void
    public sendData(event: 'mousedown' | 'mouseup' | 'keydown' | 'keyup', data: { key: number }): void
    public sendData(event: string, data: any) {
        // TODO
        if (!this.connected) {
            this.emit('warn', `attempting to send data while disconnected`)
            return
        }
        this.emit('warn', `not implemented: ${event}`)

        throw new Error('Method not implemented.')
    }

    public sendMessage(event: WebSocketEvents, payload?: WebSocketPayloads) {
        // TODO
        this.emit('debug', `sending event '${event}' ${payload ? `with payload: ` : ''}`, payload)
        throw new Error('Method not implemented.')
    }

    public async createPeer(lite: boolean, servers: RTCIceServer[]) {
        this.emit('debug', `creating peer`)
        // TODO
        throw new Error('Method not implemented.')
    }

    public async setRemoteOffer(sdp: string) {
        // TODO
        throw new Error('Method not implemented.')
    }

    public async setRemoteAnswer(sdp: string) {
        // TODO
        throw new Error('Method not implemented.')
    }

    private async onMessage(e: MessageEvent) {
        const { event, ...payload } = JSON.parse(e.data) as WebSocketMessages

        this.emit('debug', `received websocket event ${event} ${payload ? `with payload: ` : ''}`, payload)

        // TODO
        throw new Error('Method not implemented.')
    }

    private onData(e: MessageEvent) {
        this[EVENT.DATA](e.data)
    }

    private onTrack(event: RTCTrackEvent) {
        this.emit('debug', `received ${event.track.kind} track from peer: ${event.track.id}`, event)
        const stream = event.streams[0]
        if (!stream) {
            this.emit('warn', `no stream provided for track ${event.track.id}(${event.track.label})`)
            return
        }
        this[EVENT.TRACK](event)
    }

    private onError(event: Event) {
        this.emit('error', (event as ErrorEvent).error)
    }

    private onConnected() {
        if (this._timeout) {
            clearTimeout(this._timeout)
            this._timeout = undefined
        }

        if (!this.connected) {
            this.emit('warn', `onConnected called while being disconnected`)
            return
        }

        this.emit('debug', `connected`)
        this[EVENT.CONNECTED]()
    }

    private onTimeout() {
        this.emit('debug', `connection timeout`)
        if (this._timeout) {
            clearTimeout(this._timeout)
            this._timeout = undefined
        }
        this.onDisconnected(new Error('connection timeout'))
    }

    protected onDisconnected(reason?: Error) {
        this.disconnect()
        this.emit('debug', `disconnected:`, reason)
        this[EVENT.DISCONNECTED](reason)
    }

    protected [EVENT.MESSAGE](event: string, payload: any) {
        this.emit('warn', `unhandled websocket event '${event}':`, payload)
    }

    protected abstract [EVENT.RECONNECTING](): void
    protected abstract [EVENT.CONNECTING](): void
    protected abstract [EVENT.CONNECTED](): void
    protected abstract [EVENT.DISCONNECTED](reason?: Error): void
    protected abstract [EVENT.TRACK](event: RTCTrackEvent): void
    protected abstract [EVENT.DATA](data: any): void
}
