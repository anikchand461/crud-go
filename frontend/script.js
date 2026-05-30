const API_URL = "https://crudgo.onrender.com/users"

function showMessage(text, color) {

  const message = document.getElementById("message")

  message.innerText = text

  message.style.color = color
}

async function getUsers() {

  const response = await fetch(API_URL)

  const users = await response.json()

  const container = document.getElementById("users-container")

  container.innerHTML = ""

  users.forEach(user => {

    container.innerHTML += `
      <div class="user-card">

        <h2>${user.name}</h2>

        <p><strong>ID:</strong> ${user.id}</p>

        <p><strong>Age:</strong> ${user.age}</p>

        <div class="card-buttons">

          <button 
            class="edit-btn"
            onclick="fillForm(${user.id}, '${user.name}', ${user.age})"
          >
            Edit
          </button>

          <button
            class="delete-btn"
            onclick="deleteUser(${user.id})"
          >
            Delete
          </button>

        </div>

      </div>
    `
  })
}

async function addUser() {

  const id = Number(document.getElementById("id").value)

  const name = document.getElementById("name").value

  const age = Number(document.getElementById("age").value)

  const user = {
    id,
    name,
    age
  }

  const response = await fetch(API_URL, {
    method: "POST",

    headers: {
      "Content-Type": "application/json"
    },

    body: JSON.stringify(user)
  })

  const data = await response.json()

  if (!response.ok) {

    showMessage(data.error, "red")

    return
  }

  showMessage("User added successfully", "lightgreen")

  clearForm()

  getUsers()
}

async function deleteUser(id) {

  const response = await fetch(`${API_URL}/${id}`, {
    method: "DELETE"
  })

  const data = await response.json()

  if (!response.ok) {

    showMessage(data.error, "red")

    return
  }

  showMessage(data.message, "lightgreen")

  getUsers()
}

function fillForm(id, name, age) {

  document.getElementById("id").value = id

  document.getElementById("name").value = name

  document.getElementById("age").value = age
}

async function updateUser() {

  const id = document.getElementById("id").value

  const updatedUser = {
    name: document.getElementById("name").value,

    age: Number(document.getElementById("age").value)
  }

  const response = await fetch(`${API_URL}/${id}`, {
    method: "PUT",

    headers: {
      "Content-Type": "application/json"
    },

    body: JSON.stringify(updatedUser)
  })

  const data = await response.json()

  if (!response.ok) {

    showMessage(data.error, "red")

    return
  }

  showMessage("User updated successfully", "lightgreen")

  clearForm()

  getUsers()
}

function clearForm() {

  document.getElementById("id").value = ""

  document.getElementById("name").value = ""

  document.getElementById("age").value = ""
}

getUsers()