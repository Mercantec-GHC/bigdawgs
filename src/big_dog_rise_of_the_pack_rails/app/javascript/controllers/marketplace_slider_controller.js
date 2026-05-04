import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  static targets = ["range", "number", "received", "total"]
  static values = {
    rate: { type: Number, default: 1 }
  }

  connect() {
    this.sync(this.rangeTarget.value)
  }

  syncFromRange() {
    this.sync(this.rangeTarget.value)
  }

  syncFromNumber() {
    this.sync(this.numberTarget.value)
  }

  sync(value) {
    const max = Number(this.rangeTarget.max)
    const amount = Math.min(Math.max(Number(value) || 0, 0), max)
    const received = amount * this.rateValue

    this.rangeTarget.value = amount
    this.numberTarget.value = amount
    this.receivedTarget.textContent = this.format(received)
    this.totalTarget.textContent = this.format(received)
  }

  format(value) {
    return new Intl.NumberFormat().format(value)
  }
}
