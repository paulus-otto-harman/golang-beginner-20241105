
addEventListener('load', async function() {
    const token = sessionStorage.getItem("token");
    console.log(`token : ${token}`)

    const getUsers = async function() {
        const response = await fetch('/api/user/users', {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'token': token
            },
        })

        const { data } = await response.json()
        return data
    }

    const users = await getUsers()
    let userList = ''
    for await (const user of users) {
        const { id, name, is_active } = user
        userList += `<tr><td>${name}</td><td><span class="${is_active?'active':'inactive'}">${is_active?'Aktif':'Tidak Aktif'}<span></td><td><a href="/user/${id}">Detail</a></td></tr>`
    }
    const tableUsers = document.getElementById('tbl-users')
    tableUsers.innerHTML = userList
})