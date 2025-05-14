import { Controller } from "@hotwired/stimulus"

export default class extends Controller {
  static targets = [ "list", "nextPageLink" ]
  static values = {
    loadingMessage: String
  }
  static classes = [ "loading" ]

  connect() {
    this.loading = false
  }

  async loadNextPage(event) {
    event?.preventDefault()

    if (this.isLoading) return

    this.#enterLoadingState()

    const doc = await this.#fetchNextPageDocument()

    const newList = doc.querySelector(`[data-${this.identifier}-target="list"]`)

    Array.from(newList.children).forEach(child => {
      this.listTarget.appendChild(child.cloneNode(true));
    });

    const newNextPageLink = doc.querySelector(`[data-${this.identifier}-target="nextPageLink"]`)
    if (newNextPageLink) {
      this.nextPageLinkTarget.replaceWith(newNextPageLink.cloneNode(true))
    } else {
      this.nextPageLinkTarget.remove()
    }

    this.#exitLoadingState()
  }

  get isLoading() {
    return this.loading
  }

  #enterLoadingState() {
    this.loading = true

    if (this.hasLoadingMessageValue) {
      this.nextPageLinkTarget.textContent = this.loadingMessageValue
    }

    if (this.hasLoadingClass) {
      this.nextPageLinkTarget.classList.add(this.loadingClass)
    }
  }

  #exitLoadingState() {
    this.loading = false
  }

  async #fetchNextPageDocument() {
    const url = this.nextPageLinkTarget.href

    const response = await fetch(url, {
      headers: {
        "Accept": "text/html",
      }
    })

    if (response.ok) {
      return new DOMParser().parseFromString(await response.text(), "text/html")
    } else {
      console.error("Failed to fetch next page:", response.status)
      return null
    }
  }
}
