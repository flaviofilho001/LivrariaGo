const apiUrl = "http://localhost:8080";

document.getElementById("createClienteForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const data = {
    nome: document.getElementById("nome").value,
    email: document.getElementById("email").value,
    telefone: document.getElementById("telefone").value,
    cpf: document.getElementById("cpf").value,
    endereco: document.getElementById("endereco").value,
    datanascimento: document.getElementById("datanascimento").value,
    ativo: document.getElementById("ativo").checked
  };

  await fetch(`${apiUrl}/clientes/create`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data)
  });

  alert("Cliente criado com sucesso!");
  loadClientes();
});

async function loadClientes() {
  const res = await fetch(`${apiUrl}/clientes/read`);
  const clientes = await res.json();
  const tbody = document.getElementById("clientesTable");
  tbody.innerHTML = "";
  clientes.forEach(c => {
    const tr = document.createElement("tr");
    tr.innerHTML = `
      <td>${c.id}</td>
      <td>${c.nome}</td>
      <td>${c.email}</td>
      <td>${c.telefone}</td>
      <td>${c.cpf}</td>
      <td>${c.endereco}</td>
      <td>${c.datanascimento}</td>
      <td>${c.ativo ? 'Sim' : 'Não'}</td>
      <td><button onclick="deleteCliente(${c.id})">Excluir</button></td>
    `;
    tbody.appendChild(tr);
  });
}

async function deleteCliente(id) {
  await fetch(`${apiUrl}/clientes/delete?id=${id}`, { method: "DELETE" });
  alert("Cliente deletado!");
  loadClientes();
}

loadClientes();
