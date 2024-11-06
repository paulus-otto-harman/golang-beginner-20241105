
addEventListener('load', async function() {
    const token = sessionStorage.getItem("token");
    console.log(`token : ${token}`)

    const getTodos = async function() {
        const response = await fetch('/api/user/todos', {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
                'token': token
            },
        })

        const { data } = await response.json()
        return data
    }

    const todos = await getTodos()
    let todoList = ''
    for await (const todo of todos) {
        const { description, completed } = todo
        todoList += `<tr><td>${description}</td><td><span class="${completed?'done':'in-progress'}">${completed?'Done':'In Progress'}<span></td></tr>`
    }
    const tableTodos = document.getElementById('tbl-todos')
    tableTodos.innerHTML = todoList
})