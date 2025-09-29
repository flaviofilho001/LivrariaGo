const apiUrl = "http://localhost:8080";

document.getElementById("createLivroForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const data = {
    titulo: document.getElementById("titulo").value,
    autor: document.getElementById("autor").value,
    isbn: document.getElementById("isbn").value,
    preco: document.getElementById("preco").value,
    quantidade_estoque: parseInt(document.getElementById("quantidade_estoque").value),
    categoria: document.getElementById("categoria").value,
    editora: document.getElementById("editora").value,
    ano_publicacao: parseInt(document.getElementById("ano_publicacao").value),
    ativo: document.getElementById("ativo").checked
  };

  await fetch(`${apiUrl}/livros/create`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data)
  });

  alert("Livro criado com sucesso!");
  loadLivros();
});

async function loadLivros() {
  const res = await fetch(`${apiUrl}/livros/read`);
  const livros = await res.json();
  const tbody = document.getElementById("livrosTable");
  tbody.innerHTML = "";
  livros.forEach(l => {
    const tr = document.createElement("tr");
    tr.innerHTML = `
      <td>${l.id}</td>
      <td>${l.titulo}</td>
      <td>${l.autor}</td>
      <td>${l.isbn}</td>
      <td>${l.preco}</td>
      <td>${l.quantidade_estoque}</td>
      <td>${l.categoria}</td>
      <td>${l.editora}</td>
      <td>${l.ano_publicacao}</td>
      <td>${l.ativo ? 'Sim' : 'Não'}</td>
      <td><button onclick="deleteLivro(${l.id})">Excluir</button></td>
    `;
    tbody.appendChild(tr);
  });
}

async function deleteLivro(id) {
  await fetch(`${apiUrl}/livros/delete?id=${id}`, { method: "DELETE" });
  alert("Livro deletado!");
  loadLivros();
}

loadLivros();
