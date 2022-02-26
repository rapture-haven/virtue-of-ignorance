# [DRAFT]

Клиент-серверная, (!)не пир ту пир(!)

---

## Сервер 

Статика: http-сервер 

Аудио- и видео-потоки: 
* [gst-фреймворк-ом](https://en.wikipedia.org/wiki/GStreamer) создаем пайпланы для стиминга видео и аудио контекта via [RTMP](https://en.wikipedia.org/wiki/Real-Time_Messaging_Protocol) (флэшовый, медленный, траблы с мультиклиентами)
* чтобы обойти траблы с мультиклиентами, [конвертим RTMP в WebRTC](https://flashphoner.com/translyaciya-rtmp-videopotoka-iz-live-encoder-na-webrtc/?lang=ru)

Действия мыши/клавы: 
* [xOrg](https://wiki.archlinux.org/title/Xorg_(%D0%A0%D1%83%D1%81%D1%81%D0%BA%D0%B8%D0%B9)) дисплей сервер ловит действия мыши и клавы и кидает собщеньку WebRTC серверу по Data Channel, который обрабатывает соответсвующее действие

---

## Клиент 

Для работы с видео можно использовать:
```
@Ref('video') readonly _video!: HTMLVideoElement
```

У которого есть аттрибуты: 
```
video.volume
video.muted
video.stream
```

и т.д.
