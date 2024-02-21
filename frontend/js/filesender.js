function fs(){
const fileInput = document.getElementById("file-input");
const submitButton = document.getElementById("submit-button");

submitButton.addEventListener("click", async function f(){
  const file = fileInput.files[0];
  const formData = new FormData();
  formData.append("file", file);

  const response = await fetch(window.location.pathname.split('/')[1]+"/addfile/"+window.location.pathname.slice(1), {
    method: "POST",
    body: formData,
  })
    .then((response) => response.json())
    .catch((error) => console.error(error));
    alert(response)
})}
