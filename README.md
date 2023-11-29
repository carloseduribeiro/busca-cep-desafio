# Desafio Multithreading

## Curso: FullCycle - GoExpert

### Desafio:

```
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.
As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://cdn.apicep.com/file/apicep/" + cep + ".json
http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
```

### Documentações das API's utilizadas:

- ViaCEP: https://viacep.com.br
- apiCEP: https://apicep.com/api-de-consulta/

## Exemplo de execução

```shell
# faz o build do projeto e cria o binário executável 'busca-cep':
$ make build
# executa a pesquisa do cep:
$ ./busca-cep 89036-000
```