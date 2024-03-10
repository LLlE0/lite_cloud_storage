  async function req() {
    document.getElementById("name").textContent = window.location.pathname.split('/')[1] 
    try {
    const response = await fetch(window.location.pathname, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        redirect: 'manual'
    });
    
    if (response.ok) {
      const data = await response.json();
      if (data.redirect) {
          window.location.href = data.redirect;
      } else {
          const ListInst = document.getElementById('list');
          console.log(data["str"])
          console.log(Array.isArray(data["str"]))

          if (data["str"] == null) {
            const nf = document.createElement('h3')
            nf.textContent = "No files yet, go add some!"
            ListInst.appendChild(nf)
          } else {
            list = document.createElement('ul')
            ListInst.appendChild(list)
            strs = data['str']
            strs.forEach((itemUnit) => {
              const listItem = document.createElement('li');
              listItem.innerHTML=(`<a class="download-file" href="#" data-file="` + window.location.pathname.split('/')[1] + "/getfile/" + itemUnit+`">`+itemUnit+`</a>`)              
              list.appendChild(listItem);
          });
      }}
  } else {
        const errorText = await response.text();
        alert(`Login failed: ${errorText}`);
    }
} catch (error) {
    console.error('Error:', error);
}}

  $(document).ready(function() {
    // Привязываем обработчик событий к родительскому элементу
    $("#list").on("click", "a.download-file", function(event) {
      event.preventDefault(); // Отменяем стандартное поведение браузера при нажатии на ссылку
  
      var dataAddr = $(this).attr("data-file"); // Получаем адрес файла с помощью метода .attr()
      var fileName = $(this).text(); 
     $.ajax({
    url: dataAddr,
    method: 'POST',
    xhrFields: {
        responseType: 'blob'
    },
    success: function(data) {
        var a = document.createElement('a');
        var url = window.URL.createObjectURL(data);
        a.href = url;
        a.download = fileName; 
        a.click();
        window.URL.revokeObjectURL(url);
    }
});  
    });
});