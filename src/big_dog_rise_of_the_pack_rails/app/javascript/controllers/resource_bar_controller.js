import { Controller } from "@hotwired/stimulus"

// Connects to data-controller="resource-bar"
export default class extends Controller {
  connect() {
    setInterval(() => {
     this.element.reload()
    }, 10000)
  }
}
