  async function req() {
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
            strs.forEach((item) => {
              const listItem = document.createElement('li');
              listItem.textContent = item;
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