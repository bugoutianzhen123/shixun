document.getElementById("addInboundForm").addEventListener("submit", function(event) {
    event.preventDefault();

    let formData = new FormData(this);
    let data = {
        warehouseid: parseInt(formData.get("warehouseid"), 10),
        itemid: parseInt(formData.get("itemid"), 10),
        innumber: parseInt(formData.get("innumber"), 10)
    };

    fetch('http://localhost:8088/item/createinb', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': localStorage.getItem('token')
        },
        body: JSON.stringify(data)
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 200) {
                alert('Inbound record added successfully!');
                window.location.href = 'adminPanel.html';
            } else {
                alert('Failed to add inbound record: ' + data.message);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Error adding inbound record.');
        });
});

function cancelAdd() {
    window.location.href = 'adminPanel.html';
}