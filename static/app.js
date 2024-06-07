document.addEventListener('DOMContentLoaded', function () {
    fetchTodos();
});

function fetchTodos() {
    fetch('/todos')
        .then(response => response.json())
        .then(todos => {
            const todoList = document.getElementById('todo-list');
            todoList.innerHTML = '';
            todos.forEach(todo => {
                const li = document.createElement('li');
                li.textContent = `${todo.title} - ${todo.status}`;
                const deleteButton = document.createElement('button');
                deleteButton.textContent = 'Delete';
                deleteButton.onclick = () => deleteTodo(todo.id);
                li.appendChild(deleteButton);
                todoList.appendChild(li);
            });
        });
}

function addTodo() {
    const title = document.getElementById('todo-title').value;
    if (title) {
        fetch('/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ title: title, status: 'pending' })
        })
        .then(response => response.json())
        .then(() => {
            document.getElementById('todo-title').value = '';
            fetchTodos();
        });
    }
}

function deleteTodo(id) {
    fetch(`/todos/${id}`, {
        method: 'DELETE'
    })
    .then(() => fetchTodos());
}
