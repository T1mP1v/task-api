const API_URL = "http://localhost:8081/tasks";

let editingTaskId = null;

async function loadTasks(status = "") {
    let url = API_URL;
    if (status) {
        url += `?status=${status}`;
    }

    const response = await fetch(url);
    const tasks = await response.json();

    const list = document.getElementById("taskList");
    list.innerHTML = "";

    tasks.forEach(task => {
        const li = document.createElement("li");

        if (task.status === "completed") {
            li.classList.add("completed");
        }

        if (editingTaskId === task.id) {
            li.innerHTML = `
                <input type="text" id="edit-title" value="${task.title}">
                <input type="text" id="edit-desc" value="${task.description || ""}">
                <div class="actions">
                    <button onclick="saveTask(${task.id})">💾</button>
                    <button onclick="cancelEdit()">✕</button>
                </div>
            `;
        } else {
            li.innerHTML = `
                <span>
                    <strong>${task.title}</strong><br>
                    ${task.description || ""}
                </span>
                <div class="actions">
                    <button onclick="startEdit(${task.id})">✎</button>
                    <button onclick="toggleStatus(${task.id}, '${task.status}')">✓</button>
                    <button onclick="deleteTask(${task.id})">🗑</button>
                </div>
            `;
        }

        list.appendChild(li);
    });
}

function startEdit(id) {
    editingTaskId = id;
    loadTasks();
}

function cancelEdit() {
    editingTaskId = null;
    loadTasks();
}

async function saveTask(id) {
    const title = document.getElementById("edit-title").value;
    const description = document.getElementById("edit-desc").value;

    if (!title) {
        alert("Название не может быть пустым");
        return;
    }

    await fetch(`${API_URL}/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            title: title,
            description: description
        })
    });

    editingTaskId = null;
    loadTasks();
}

async function createTask() {
    const title = document.getElementById("title").value;
    const description = document.getElementById("description").value;

    if (!title) {
        alert("Введите название задачи");
        return;
    }

    await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            title: title,
            description: description,
            status: "new"
        })
    });

    document.getElementById("title").value = "";
    document.getElementById("description").value = "";

    loadTasks();
}

async function deleteTask(id) {
    await fetch(`${API_URL}/${id}`, {
        method: "DELETE"
    });
    loadTasks();
}

async function toggleStatus(id, currentStatus) {
    const newStatus = currentStatus === "new" ? "completed" : "new";

    await fetch(`${API_URL}/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            status: newStatus
        })
    });

    loadTasks();
}

loadTasks();
