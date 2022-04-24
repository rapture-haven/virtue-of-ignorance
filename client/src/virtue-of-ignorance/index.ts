import Vue from "vue";
import EventEmitter from "eventemitter3";
import { BaseClient, BaseEvents } from "./base";
import { EVENT } from "./events";
import { accessor } from "@/store";

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
} from "./messages";

type VirtueEvents = BaseEvents;

export class VirtueClient
  extends BaseClient
  implements EventEmitter<VirtueEvents>
{
  private $vue!: Vue;
  private $accessor!: typeof accessor;
  private url!: string;

  init(vue: Vue) {
    const url =
      process.env.NODE_ENV === "development"
        ? `ws://${location.host.split(":")[0]}:${
            process.env.VUE_APP_SERVER_PORT
          }/ws`
        : location.protocol.replace(/^http/, "ws") +
          "//" +
          location.host +
          location.pathname.replace(/\/$/, "") +
          "/ws";

    this.initWithURL(vue, url);
  }

  initWithURL(vue: Vue, url: string) {
    this.$vue = vue;
    this.$accessor = vue.$accessor;
    this.url = url;
  }

  private cleanup() {
    this.$accessor.setConnected(false)
    this.$accessor.remote.reset()
    this.$accessor.user.reset()
    this.$accessor.video.reset()
  }
  
  login(password: string, displayname: string) {
    this.connect(this.url, password, displayname);
  }

  logout() {
    this.disconnect();
    this.cleanup()
    this.$vue.$swal({
      title: this.$vue.$t('connection.logged_out'),
      icon: 'info',
      confirmButtonText: this.$vue.$t('connection.button_confirm') as string,
    })
  }

  protected [EVENT.RECONNECTING](): void {
    this.$vue.$notify({
      group: 'neko',
      type: 'warning',
      title: this.$vue.$t('connection.reconnecting') as string,
      duration: 5000,
      speed: 1000,
    })
  }

  protected [EVENT.CONNECTING](): void {
    this.$accessor.setConnnecting();
  }

  protected [EVENT.CONNECTED](): void {
    this.$accessor.user.setMember(this.id)
    this.$accessor.setConnected(true);

    this.$vue.$notify({
      group: "virtue",
      clean: true,
    });

    this.$vue.$notify({
      group: "virtue",
      type: "success",
      title: "Подключено" as string,
      duration: 5000,
      speed: 1000,
    });
  }

  protected [EVENT.DISCONNECTED](reason?: Error): void {
    this.cleanup()

    this.$vue.$notify({
      group: 'neko',
      type: 'error',
      title: this.$vue.$t('connection.disconnected') as string,
      text: reason ? reason.message : undefined,
      duration: 5000,
      speed: 1000,
    })
  }

  protected [EVENT.TRACK](event: RTCTrackEvent): void {
    throw new Error("Method not implemented.");
  }
  
  protected [EVENT.DATA](data: any): void { }
}
