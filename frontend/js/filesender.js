function fs() { 
const fileInput = document.getElementById("file-input");
const submitButton = document.getElementById("submit-button");

submitButton.addEventListener("click", async function f(event){
  event.preventDefault(); // Отменяем стандартное поведение браузера при нажатии на кнопку

  const fileList = fileInput.files;
  const file = fileList[0];

  if (file) {
    const formData = new FormData();
    formData.append("file", file);

    const response = await fetch(window.location.pathname.split('/')[1]+"/addfile/"+window.location.pathname.slice(1), {
      method: "POST",
      body: formData,
    })
    .then((response) => response.json())
    .catch((error) => console.error(error));
    alert(response)
  } else {
    alert("Please select the file first!")
  }
})}