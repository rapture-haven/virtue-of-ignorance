import Vue from 'vue'
import EventEmitter from 'eventemitter3'
import { BaseClient, BaseEvents } from './base'
import { EVENT } from './events'
import { accessor } from '@/store'

import {
    SystemMessagePayload,
    SignalProvidePayload,
    MemberListPayload,
    MemberDisconnectPayload,
    MemberPayload,
    ControlPayload,
    ControlTargetPayload,
    ControlClipboardPayload,
    ScreenConfigurationsPayload,
    ScreenResolutionPayload,
    BroadcastStatusPayload,
    AdminPayload,
    AdminTargetPayload,
    AdminLockMessage,
    SystemInitPayload,
    AdminLockResource,
} from './messages'

interface VirtueEvents extends BaseEvents { }

export class VirtueClient extends BaseClient implements EventEmitter<VirtueEvents> {
    private $vue!: Vue
    private $accessor!: typeof accessor
    private url!: string

    init(vue: Vue) {
        const url =
            process.env.NODE_ENV === 'development'
                ? `ws://${location.host.split(':')[0]}:${process.env.VUE_APP_SERVER_PORT}/ws`
                : location.protocol.replace(/^http/, 'ws') + '//' + location.host + location.pathname.replace(/\/$/, '') + '/ws'

        this.initWithURL(vue, url)
    }

    initWithURL(vue: Vue, url: string) {
        this.$vue = vue
        this.$accessor = vue.$accessor
        this.url = url
    }

    login(password: string, displayname: string) {
        this.connect(this.url, password, displayname)
    }

    logout() {
        this.disconnect()
    }

    protected [EVENT.RECONNECTING](): void {
        throw new Error('Method not implemented.')
    }
    protected [EVENT.CONNECTING](): void {
        this.$accessor.setConnnecting()
    }
    protected [EVENT.CONNECTED](): void {
        // this.$accessor.user.setMember(this.id)
        this.$accessor.setConnected(true)

        this.$vue.$notify({
            group: 'virtue',
            clean: true,
        })

        this.$vue.$notify({ 
            group: 'virtue',
            type: 'success',
            title: "Подключено" as string,
            duration: 5000,
            speed: 1000,
        })
    }
    protected [EVENT.DISCONNECTED](reason?: Error): void {
        throw new Error('Method not implemented.')
    }
    protected [EVENT.TRACK](event: RTCTrackEvent): void {
        throw new Error('Method not implemented.')
    }
    protected [EVENT.DATA](data: any): void {
        throw new Error('Method not implemented.')
    }
};
