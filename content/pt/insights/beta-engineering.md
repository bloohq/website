---
title: Por que Blue tem um Beta Aberto
category: "Engineering"
description: Saiba por que nosso sistema de gerenciamento de projetos tem um beta aberto contínuo.
date: 2024-08-03
---

Muitas startups B2B SaaS lançam em Beta, e por boas razões. Faz parte do tradicional lema do Vale do Silício de *"mover-se rápido e quebrar coisas"*.

Colocar um selo "beta" em um produto reduz as expectativas.

Algo está quebrado? Bem, é apenas um beta.

O sistema está lento? Bem, é apenas um beta.

[A documentação](https://blue.cc/docs) é inexistente? Bem... você entendeu.

E isso é *realmente* uma coisa boa. Reid Hoffman, o fundador do LinkedIN, disse a famosa frase:

> Se você não está envergonhado pela primeira versão do seu produto, você lançou tarde demais.

E o selo beta também é bom para os clientes. Ajuda-os a se auto-selecionar.

Os clientes que experimentam produtos beta são aqueles nos estágios iniciais do Ciclo de Vida de Adoção de Tecnologia, também conhecido como Curva de Adoção de Produto.

O Ciclo de Vida de Adoção de Tecnologia é tipicamente dividido em cinco segmentos principais:

1. Inovadores
2. Adotantes Iniciais
3. Maioria Inicial
4. Maioria Tardia
5. Retardatários

![](/insights/technology-adoption-lifecycle-graph.png)

No entanto, eventualmente o produto tem que amadurecer, e os clientes esperam um produto estável e funcional. Eles não querem acesso a um ambiente "beta" onde as coisas quebram.

Ou querem?

*Essa* é a pergunta que nos fizemos.

Acreditamos que nos fizemos essa pergunta devido à natureza de como Blue foi inicialmente construído. [Blue começou como um desdobramento de uma agência de design movimentada](/insights/agency-success-playbook), e então trabalhamos *dentro* do escritório de uma empresa que estava ativamente usando Blue para gerenciar todos os seus projetos.

Isso significa que por anos, pudemos observar como *seres humanos reais* — sentados bem ao nosso lado! — usavam Blue em suas vidas diárias.

E porque eles usavam Blue desde os primeiros dias, essa equipe sempre usou Blue Beta!

E então foi natural para nós permitir que todos os nossos outros clientes também o usassem.

**E é por isso que não temos uma equipe de testes dedicada.**

Isso mesmo.

Ninguém em Blue tem a *única* responsabilidade de garantir que nossa plataforma esteja funcionando bem e estável.

Isso acontece por várias razões.

A primeira é uma base de custos mais baixa.

Não ter uma equipe de testes em tempo integral reduz significativamente nossos custos, e podemos repassar essas economias aos nossos clientes com os preços mais baixos do setor.

Para colocar isso em perspectiva, oferecemos conjuntos de recursos de nível empresarial que nossa concorrência cobra $30-$55/usuário/mês por apenas $7/mês.

Isso não acontece por acidente, é *por design*.

No entanto, não é uma boa estratégia vender um produto mais barato se ele não funciona.

Então a *verdadeira questão é*, como conseguimos criar uma plataforma estável que milhares de clientes podem usar sem uma equipe de testes dedicada?

É claro que nossa abordagem de ter um Beta aberto é crucial para isso, mas antes de mergulharmos nisso, queremos tocar na responsabilidade do desenvolvedor.

Tomamos a decisão inicial em Blue de que nunca dividiríamos as responsabilidades para tecnologias de front-end e back-end. Contrataríamos ou treinaríamos apenas desenvolvedores full stack.

A razão pela qual tomamos essa decisão foi para garantir que um desenvolvedor seria totalmente dono do recurso em que estava trabalhando. Então não haveria nada da mentalidade de *"jogar o problema por cima do muro"* que às vezes você encontra quando há responsabilidades conjuntas pelos recursos.

E isso se estende ao teste do recurso, ao entendimento dos casos de uso e solicitações do cliente, e à leitura e comentários nas especificações.

Em outras palavras, cada desenvolvedor constrói um entendimento profundo e intuitivo do recurso que está construindo.

Ok, vamos agora falar sobre nosso beta aberto.

Quando dizemos que é "aberto" — queremos dizer isso. Qualquer cliente pode experimentá-lo simplesmente adicionando "beta" na frente da url da nossa aplicação web.

Então "app.blue.cc" se torna "beta.app.blue.cc"

Quando fazem isso, eles podem ver seus dados usuais, já que os ambientes Beta e Produção compartilham o mesmo banco de dados, mas também poderão ver novos recursos.

Os clientes podem trabalhar facilmente mesmo que tenham alguns membros da equipe em Produção e alguns curiosos em Beta.

Normalmente temos algumas centenas de clientes usando Beta a qualquer momento, e publicamos prévias de recursos em nossos fóruns da comunidade para que possam conferir as novidades e experimentá-las.

E este é o ponto: temos *várias centenas* de testadores!

Todos esses clientes experimentarão recursos em seus fluxos de trabalho e serão bastante vocais se algo não estiver certo, porque eles *já estão* implementando o recurso dentro de seus negócios!

O feedback mais comum são pequenas mas muito úteis mudanças que abordam casos extremos que não consideramos.

Deixamos novos recursos em Beta entre 2-4 semanas. Sempre que sentimos que estão estáveis, então lançamos para produção.

Também temos a capacidade de contornar o Beta se necessário, usando uma flag de fast-track. Isso é normalmente feito para correções de bugs que não queremos segurar por 2-4 semanas antes de enviar para produção.

O resultado?

Fazer push para produção parece... bem entediante! Como nada — simplesmente não é um grande negócio para nós.

E isso significa que suaviza nosso cronograma de lançamento, que é o que nos permitiu [lançar recursos mensalmente como um relógio nos últimos seis anos.](/changelog).

No entanto, como qualquer escolha, existem alguns trade-offs.

O suporte ao cliente é ligeiramente mais complexo, pois temos que dar suporte a clientes em duas versões de nossa plataforma. Às vezes isso pode causar confusão aos clientes que têm membros da equipe usando duas versões diferentes.

Outro ponto problemático é que essa abordagem às vezes pode desacelerar o cronograma geral de lançamento para produção. Isso é especialmente verdadeiro para recursos maiores que podem ficar "presos" em Beta se houver outro recurso relacionado que está tendo problemas e precisa de mais trabalho.

Mas no geral, achamos que esses trade-offs valem os benefícios de uma base de custos mais baixa e maior engajamento do cliente.

Somos uma das poucas empresas de software a abraçar essa abordagem, mas agora é uma parte fundamental de nossa abordagem de desenvolvimento de produto.
