import { Controller } from "@hotwired/stimulus"

// Connects to data-controller="construction"
export default class extends Controller {
  static values = {
    endtime: { type: Number, default: 0 },
  }
  connect() {
    setInterval(() => {
      this.endtimeValue -= 1
      if (this.endtimeValue <= 0) {
        this.endtimeValue = 0
      }
      this.element.textContent = `Time left: ${Math.floor(this.endtimeValue)} seconds`
    }, 1000);
  }
}
