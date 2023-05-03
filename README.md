# Exercicio docker go fullcycle

```bash
# Proposito do exercicio 

Creio que o proposito seria criar o multi stage para deixar mais leve,
eu nao encontrei um hello word ! sem criar o servidor com o go, 
por isso creio que minha imagem passou os 2MB, mas consegui criar o
multi stage deixando a imagem de 950MB com 25MB

# Funcionamento da imagem

Basta rodar docker run -p 8080:8080 -t nomeimagem .
E rodar no navegador http://localhost:8080/FullCycleRocks

```