
document.getElementById('btn-save').addEventListener('click', async function (e) {
    e.preventDefault()
    if (document.getElementById('frm-registration').reportValidity()) {
        const form = document.querySelector('form');
        const params = new FormData(form);
        const pairs = {};
        for (const [name, value] of params) {
            pairs[name] = value;
        }

        // '/api/register'
        const response = await fetch('https://paulus.free.beeceptor.com/register', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(pairs),

        })
        const { data : {id} } = await response.json();

        sessionStorage.clear()
        sessionStorage.setItem("token", id);
        window.location.href = '/users'
    }

})