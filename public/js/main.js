(() => {
  // <stdin>
  console.log("This site was made by Natalia with <3.");
  document.addEventListener("scroll", function() {
    const logo = document.querySelector(".logo");
    if (window.scrollY > 100) {
      logo.style.opacity = "0";
      logo.style.transition = "opacity 2s ease";
    } else {
      logo.style.opacity = "1";
    }
  });
})();
