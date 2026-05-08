import { Controller } from "@hotwired/stimulus"

// Connects to data-controller="marketplace-refresh"
export default class extends Controller {
  connect() {
    let Timer = 300
    this.element.textContent = `Next market refresh in: ${Timer} seconds`
    const interval = setInterval(() => {
      Timer -= 1
      this.element.textContent = `Next market refresh in: ${Timer} seconds`

      if (Timer <= 0) {
        clearInterval(interval)   // 👈 stopper intervallet
        location.reload()
      }
    }, 1000)
  }

  
}
