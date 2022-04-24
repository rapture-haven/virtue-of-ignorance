/* eslint-disable @typescript-eslint/ban-ts-comment */
export function elementRequestFullscreen(el: HTMLElement) {
  if (typeof el.requestFullscreen === "function") {
    el.requestFullscreen();
    //@ts-ignore
  } else if (typeof el.webkitRequestFullscreen === "function") {
    //@ts-ignore
    el.webkitRequestFullscreen();
    //@ts-ignore
  } else if (typeof el.webkitEnterFullscreen === "function") {
    //@ts-ignore
    el.webkitEnterFullscreen();
    //@ts-ignore
  } else if (typeof el.mozRequestFullScreen === "function") {
    //@ts-ignore
    el.mozRequestFullScreen();
    //@ts-ignore
  } else if (typeof el.msRequestFullScreen === "function") {
    //@ts-ignore
    el.msRequestFullScreen();
  } else {
    return false;
  }
  return true;
}

export function isFullscreen(): boolean {
  return (
    document.fullscreenElement ||
    //@ts-ignore
    document.msFullscreenElement ||
    //@ts-ignore
    document.mozFullScreenElement ||
    //@ts-ignore
    document.webkitFullscreenElement
  );
}

export function onFullscreenChange(el: HTMLElement, fn: () => void) {
  if (el.onfullscreenchange === null) {
    el.onfullscreenchange = fn;
    //@ts-ignore
  } else if (el.onmsfullscreenchange === null) {
    //@ts-ignore
    el.onmsfullscreenchange = fn;
    //@ts-ignore
  } else if (el.onmozfullscreenchange === null) {
    //@ts-ignore
    el.onmozfullscreenchange = fn;
    //@ts-ignore
  } else if (el.onwebkitfullscreenchange === null) {
    //@ts-ignore
    el.onwebkitfullscreenchange = fn;
  }
}
/* eslint-enable @typescript-eslint/ban-ts-comment */
