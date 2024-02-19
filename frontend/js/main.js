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
          if (Array.isArray(data["str"]) && data["str"].length === 0) {
            const nf = document.createElement('h3')
            nf.textContent = "No files yet, go add some!"
            nf.text="No files!"
            ListInst.appendChild(nf)
            console.log(1)
          } else {
            console.log(2)
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