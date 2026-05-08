import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  static targets = ["popup"]

  connect() {
    this.positionAfterLoad = this.positionAfterLoad.bind(this)
    this.popupTarget.addEventListener("turbo:frame-load", this.positionAfterLoad)
  }

  disconnect() {
    this.popupTarget.removeEventListener("turbo:frame-load", this.positionAfterLoad)
  }

  move(event) {
    this.mouseX = event.clientX
    this.mouseY = event.clientY
  }

  positionAfterLoad() {
    const frame = this.popupTarget
    const popup = frame.firstElementChild

    if (!popup) return

    const popupWidth = popup.offsetWidth
    const popupHeight = popup.offsetHeight
    const gap = 12

    let x = this.mouseX + gap
    let y = this.mouseY + gap

    if (this.mouseX > window.innerWidth / 2) {
      x = this.mouseX - popupWidth - gap
    }

    if (this.mouseY > window.innerHeight / 2) {
      y = this.mouseY - popupHeight - gap
    }

    frame.style.left = `${x}px`
    frame.style.top = `${y}px`
  }

  close(event){
    this.popupTarget.innerHTML = ""
  }
}