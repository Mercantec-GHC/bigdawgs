import "@hotwired/turbo-rails"
import "controllers"

document.addEventListener("turbo:load", () => {
  console.log("Turbo loaded, initializing toasts...")
  const toastElList = document.querySelectorAll(".toast")
  toastElList.forEach((toastEl) => {
    const toast = new bootstrap.Toast(toastEl)
    toast.show()
  })
})


