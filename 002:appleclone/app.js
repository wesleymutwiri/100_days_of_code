let menu = document.querySelectorAll(".item");

let navbarToggle = document.getElementById("toggle");

navbarToggle.addEventListener("click", function () {
  menu.forEach((menu) => menu.classList.toggle("active"));
});
