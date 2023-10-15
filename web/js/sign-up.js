document.getElementById("myForm").addEventListener("submit", function (event) {
    event.preventDefault();

    const form = event.target;
    const formData = new FormData(form);

    const data = {};
    formData.forEach((value, key) => {
        data[key] = value;
    });

    const url = "/auth/sign-up";
    const options = {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json"
        }
    };

    fetch(url, options)
        .then(response => response.json())
        .then(data => {
            // Обработка успешного ответа от сервера (data)
            console.log(data);
        })
        .catch(error => {
            // Обработка ошибок
            console.error(error);
        });
});