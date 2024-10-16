document.addEventListener('DOMContentLoaded', function () {
    loadUsers();
    loadPlans();
    loadPayments();

    // Criação de usuário
    document.getElementById('userForm').addEventListener('submit', createUser);

    // Criação de plano
    document.getElementById('planForm').addEventListener('submit', createPlan);

    // Criação de pagamento
    document.getElementById('paymentForm').addEventListener('submit', createPayment);

    // Carrega a tabela de usuários
    function loadUsers() {
        fetch('/users/read')
            .then(response => response.json())
            .then(data => {
                console.log(data); // Para verificar a estrutura dos dados
                const userTableBody = document.querySelector('#userTable tbody');
                userTableBody.innerHTML = ''; // Limpa a tabela antes de adicionar novos dados
                
                data.forEach(user => {
                    const row = document.createElement('tr');
                    row.innerHTML = `
                        <td>${user.Id !== undefined ? user.Id : ''}</td> <!-- Corrigido para usar "Id" -->
                        <td>${user.Name !== undefined ? user.Name : ''}</td> <!-- Corrigido para usar "Name" -->
                        <td>${user.Email !== undefined ? user.Email : ''}</td> <!-- Corrigido para usar "Email" -->
                        <td>${user.Age !== undefined ? user.Age : ''}</td> <!-- Corrigido para usar "Age" -->
                        <td>${user.Plan && user.Plan.Name ? user.Plan.Name : ''}</td> <!-- Corrigido para usar "Plan.Name" -->
                        <td>${user.Status !== undefined ? (user.Status ? 'Ativo' : 'Inativo') : ''}</td> <!-- Corrigido para usar "Status" -->
                        <td>
                            <button data-id="${user.Id}" class="edit-button">Editar</button>
                            <button data-id="${user.Id}" class="delete-button">Excluir</button>
                        </td>
                    `;
                    userTableBody.appendChild(row);
                });
    
                // Adicionar evento de clique para editar e excluir
                document.querySelectorAll('.edit-button').forEach(button => {
                    button.addEventListener('click', editUser);
                });
    
                document.querySelectorAll('.delete-button').forEach(button => {
                    button.addEventListener('click', deleteUser);
                });
            })
            .catch(error => {
                console.error('Erro ao carregar usuários:', error);
            });
    }
    
    

    // Carrega a tabela de planos
    function loadPlans() {
        fetch('/plans/read')
            .then(response => response.json())
            .then(data => {
                console.log(data); // Para verificar a estrutura dos dados
                const planTableBody = document.querySelector('#planTable tbody');
                planTableBody.innerHTML = ''; // Limpa a tabela antes de adicionar novos dados
                
                data.forEach(plan => {
                    const row = document.createElement('tr');
                    row.innerHTML = `
                        <td>${plan.Id !== undefined ? plan.Id : ''}</td> <!-- Corrigido para usar "Id" -->
                        <td>${plan.Name !== undefined ? plan.Name : ''}</td> <!-- Corrigido para usar "Name" -->
                        <td>${plan.Price !== undefined ? plan.Price : ''}</td> <!-- Corrigido para usar "Price" -->
                        <td>
                            <button data-id="${plan.Id}" class="edit-button">Editar</button>
                            <button data-id="${plan.Id}" class="delete-button">Excluir</button>
                        </td>
                    `;
                    planTableBody.appendChild(row);
                });
    
                // Adicionar evento de clique para editar e excluir
                document.querySelectorAll('.edit-button').forEach(button => {
                    button.addEventListener('click', editPlan);
                });
    
                document.querySelectorAll('.delete-button').forEach(button => {
                    button.addEventListener('click', deletePlan);
                });
            })
            .catch(error => {
                console.error('Erro ao carregar planos:', error);
            });
    }
    

    // Carrega a tabela de pagamentos
    function loadPayments() {
        fetch('/payments/read')
            .then(response => response.json())
            .then(data => {
                console.log(data); // Para verificar a estrutura dos dados
                const paymentTableBody = document.querySelector('#paymentTable tbody');
                paymentTableBody.innerHTML = ''; // Limpa a tabela antes de adicionar novos dados
                
                data.forEach(payment => {
                    const row = document.createElement('tr');
                    row.innerHTML = `
                        <td>${payment.Id !== undefined ? payment.Id : ''}</td> <!-- Corrigido para usar "Id" -->
                        <td>${payment.UserName !== undefined ? payment.UserName : ''}</td> <!-- Corrigido para usar "UserName" -->
                        <td>${payment.Month !== undefined ? payment.Month : ''}</td> <!-- Corrigido para usar "Month" -->
                        <td>${payment.Status !== undefined ? (payment.Status ? 'Pago' : 'Não Pago') : ''}</td> <!-- Corrigido para usar "Status" -->
                        <td>${payment.PaymentDate !== undefined ? payment.PaymentDate : ''}</td> <!-- Corrigido para usar "PaymentDate" -->
                        <td>
                            <button data-id="${payment.Id}" class="edit-button">Editar</button>
                            <button data-id="${payment.Id}" class="delete-button">Excluir</button>
                        </td>
                    `;
                    paymentTableBody.appendChild(row);
                });
    
                // Adicionar evento de clique para editar e excluir
                document.querySelectorAll('.edit-button').forEach(button => {
                    button.addEventListener('click', editPayment);
                });
    
                document.querySelectorAll('.delete-button').forEach(button => {
                    button.addEventListener('click', deletePayment);
                });
            })
            .catch(error => {
                console.error('Erro ao carregar pagamentos:', error);
            });
    }
    

    // Função para criar usuário
    function createUser(e) {
        e.preventDefault();
        const name = document.getElementById('name').value;
        const email = document.getElementById('email').value;
        const age = document.getElementById('age').value;
        const planId = document.getElementById('planId').value;
        const status = document.getElementById('status').checked;

        fetch('/users/create', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                name: name,
                email: email,
                age: age,
                plan: { id: planId },
                status: status
            })
        })
        .then(response => {
            if (response.ok) {
                loadUsers(); // Recarrega a tabela de usuários
                clearUserForm(); // Limpa o formulário após a criação
            }
        });
    }

    // Função para criar plano
    function createPlan(e) {
        e.preventDefault();
        const planName = document.getElementById('planName').value;
        const planPrice = document.getElementById('planPrice').value;

        fetch('/plans/create', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                name: planName,
                price: planPrice
            })
        })
        .then(response => {
            if (response.ok) {
                loadPlans(); // Recarrega a tabela de planos
                clearPlanForm(); // Limpa o formulário após a criação
            }
        });
    }

    // Função para criar pagamento
    function createPayment(e) {
        e.preventDefault();
        const userId = document.getElementById('userId').value;
        const month = document.getElementById('month').value;
        const paymentStatus = document.getElementById('paymentStatus').checked;
    
        // Verifica se o mês está definido antes de fazer a requisição
        if (!month) {
            alert("Por favor, selecione um mês.");
            return;
        }
    
        fetch('/payments/create', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                userId: userId,
                month: month,
                status: paymentStatus
            })
        })
        .then(response => {
            if (response.ok) {
                loadPayments(); // Recarrega a tabela de pagamentos
                clearPaymentForm(); // Limpa o formulário após a criação
            }
        });
    }
    

    // Função para editar usuário
    function editUser(e) {
        const userId = e.target.getAttribute('data-id');
        fetch(`/users/read?id=${userId}`)
            .then(response => response.json())
            .then(user => {
                document.getElementById('name').value = user.name;
                document.getElementById('email').value = user.email;
                document.getElementById('age').value = user.age;
                document.getElementById('planId').value = user.plan ? user.plan.id : ''; // Selecionar o plano
                document.getElementById('status').checked = user.status;
    
                // Verifica se o campo de mês existe e define um valor padrão ou vazio
                const monthInput = document.getElementById('month');
                if (user.month) {
                    monthInput.value = user.month; // Supondo que você tenha um mês associado ao usuário
                } else {
                    monthInput.value = ''; // Se não houver, define como vazio
                }
    
                const userForm = document.getElementById('userForm');
                userForm.removeEventListener('submit', createUser); // Remove o evento anterior
                userForm.addEventListener('submit', function (e) {
                    e.preventDefault();
                    updateUser(userId); // Chama a função para atualizar o usuário
                });
            });
    }
    

    // Função para atualizar usuário
    function updateUser(userId) {
        const name = document.getElementById('name').value;
        const email = document.getElementById('email').value;
        const age = document.getElementById('age').value;
        const planId = document.getElementById('planId').value;
        const status = document.getElementById('status').checked;

        fetch(`/users/update?id=${userId}`, {
            method: 'PUT', // ou 'PATCH', dependendo de como você configurou sua API
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: name,
                email: email,
                age: age,
                plan: { id: planId },
                status: status
            })
        })
        .then(response => {
            if (response.ok) {
                loadUsers(); // Recarrega a tabela de usuários
                clearUserForm(); // Limpa o formulário após a atualização
            } else {
                console.error("Erro ao atualizar usuário.");
            }
        });
    }

    // Função para excluir usuário
    function deleteUser(e) {
        const userId = e.target.getAttribute('data-id');
        fetch(`/users/delete?id=${userId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                loadUsers(); // Recarrega a tabela de usuários
            } else {
                console.error("Erro ao excluir usuário.");
            }
        });
    }

    // Função para excluir plano
    function deletePlan(e) {
        const planId = e.target.getAttribute('data-id');
        fetch(`/plans/delete?id=${planId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                loadPlans(); // Recarrega a tabela de planos
            } else {
                console.error("Erro ao excluir plano.");
            }
        });
    }

    // Função para excluir pagamento
    function deletePayment(e) {
        const paymentId = e.target.getAttribute('data-id');
        fetch(`/payments/delete?id=${paymentId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                loadPayments(); // Recarrega a tabela de pagamentos
            } else {
                console.error("Erro ao excluir pagamento.");
            }
        });
    }

    // Funções para limpar formulários
    function clearUserForm() {
        document.getElementById('name').value = '';
        document.getElementById('email').value = '';
        document.getElementById('age').value = '';
        document.getElementById('planId').value = '';
        document.getElementById('status').checked = false;
    }

    function clearPlanForm() {
        document.getElementById('planName').value = '';
        document.getElementById('planPrice').value = '';
    }

    function clearPaymentForm() {
        document.getElementById('userId').value = '';
        document.getElementById('month').value = '';
        document.getElementById('paymentStatus').checked = false;
    }
});

