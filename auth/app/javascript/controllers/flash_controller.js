import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  static values = {
    timeout: { type: Number, default: 2000 },
  }

  connect() {
    this.timeout = setTimeout(() => {
      this.remove()
    }, this.timeoutValue)
  }

  remove() {
    this.element.remove()
  }
}
