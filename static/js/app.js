// static/js/app.js

document.addEventListener('DOMContentLoaded', () => {
    const userForm = document.getElementById('user-form');
    const cancelBtn = document.getElementById('cancel-btn');
    const usersTableBody = document.querySelector('#users-table tbody');

    let editMode = false;
    let editUserId = null;

    // Função para buscar e exibir usuários
    const fetchUsers = async () => {
        try {
            const response = await fetch('/api/users');
            const data = await response.json();
            console.log('Dados da API:', data); // Verifique se a estrutura dos dados está correta
            populateTable(data.data);
        } catch (error) {
            console.error('Erro ao buscar usuários:', error);
        }
    };    

    // Função para popular a tabela com usuários
    const populateTable = (users) => {
        usersTableBody.innerHTML = '';
        users.forEach(user => {
            const row = document.createElement('tr');

            row.innerHTML = `
                <td>${user.ID}</td>
                <td>${user.Name}</td>
                <td>${user.Email}</td>
                <td>${user.Age}</td>
                <td>
                    <button class="action-btn edit-btn" data-id="${user.ID}">Editar</button>
                    <button class="action-btn delete-btn" data-id="${user.ID}">Deletar</button>
                </td>
            `;
            usersTableBody.appendChild(row);
        });
    };

    // Função para adicionar ou editar usuário
    userForm.addEventListener('submit', async (e) => {
        e.preventDefault();

        const name = document.getElementById('name').value;
        const email = document.getElementById('email').value;
        const age = parseInt(document.getElementById('age').value, 10);

        const userData = { Name: name, Email: email, Age: age };

        try {
            let response;
            if (editMode) {
                response = await fetch(`/api/users/${editUserId}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(userData)
                });
            } else {
                response = await fetch('/api/users', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(userData)
                });
            }

            if (!response.ok) {
                const errorData = await response.json();
                alert(`Erro: ${errorData.error}`);
                return;
            }

            userForm.reset();
            editMode = false;
            editUserId = null;
            cancelBtn.style.display = 'none';
            fetchUsers();
        } catch (error) {
            console.error('Erro ao salvar usuário:', error);
        }
    });

    // Função para editar usuário
    usersTableBody.addEventListener('click', async (e) => {
        if (e.target.classList.contains('edit-btn')) {
            const userId = e.target.getAttribute('data-id');
            try {
                const response = await fetch(`/api/users/${userId}`);
                const data = await response.json();
                const user = data.data;

                document.getElementById('user-id').value = user.ID;
                document.getElementById('name').value = user.Name;
                document.getElementById('email').value = user.Email;
                document.getElementById('age').value = user.Age;

                editMode = true;
                editUserId = user.ID;
                cancelBtn.style.display = 'inline-block';
            } catch (error) {
                console.error('Erro ao buscar usuário:', error);
            }
        }

        // Função para deletar usuário
        if (e.target.classList.contains('delete-btn')) {
            const userId = e.target.getAttribute('data-id');
            if (confirm('Tem certeza que deseja deletar este usuário?')) {
                try {
                    const response = await fetch(`/api/users/${userId}`, {
                        method: 'DELETE'
                    });

                    if (!response.ok) {
                        const errorData = await response.json();
                        alert(`Erro: ${errorData.error}`);
                        return;
                    }

                    fetchUsers();
                } catch (error) {
                    console.error('Erro ao deletar usuário:', error);
                }
            }
        }
    });

    // Função para cancelar edição
    cancelBtn.addEventListener('click', () => {
        userForm.reset();
        editMode = false;
        editUserId = null;
        cancelBtn.style.display = 'none';
    });

    // Inicializar listagem de usuários
    fetchUsers();
});
