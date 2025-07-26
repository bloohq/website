---
title: Come creare automazioni email personalizzate
description: Le notifiche email personalizzate sono una funzionalità incredibilmente potente in Blue che può aiutare a mantenere il lavoro in movimento e garantire che la comunicazione sia in modalità automatica.
category: "Product Updates"
---


Le automazioni email in Blue sono una [potente automazione per la gestione dei progetti](/platform/features/automations) per semplificare la comunicazione, garantire [ottimo lavoro di squadra](/insights/great-teamwork) e mantenere i progetti in avanti. Sfruttando i dati memorizzati nei tuoi record, puoi inviare automaticamente email personalizzate quando si verificano determinati trigger, come la creazione di un nuovo record o una scadenza di un'attività.

In questo articolo, esploreremo come impostare e utilizzare le automazioni email in Blue.

## Impostazione delle Automazioni Email

Creare un'automazione email in Blue è un processo semplice. Prima di tutto, seleziona il trigger che avvierà l'email automatizzata. Alcuni trigger comuni includono:

- Un nuovo record viene creato
- Un tag viene aggiunto a un record
- Un record viene spostato in un'altra lista

Successivamente, configura i dettagli dell'email, inclusi:

- Nome del mittente e indirizzo di risposta
- Indirizzo del destinatario (può essere statico o estratto dinamicamente da un campo email personalizzato)
- Indirizzi CC o BCC (opzionale)

![](/insights/email-automations-image.webp)

Uno dei principali vantaggi delle automazioni email in Blue è la possibilità di personalizzare il contenuto utilizzando i merge tag. Quando personalizzi l'oggetto e il corpo dell'email, puoi inserire merge tag che fanno riferimento a dati specifici del record, come il nome del record o i valori dei campi personalizzati. Basta utilizzare la sintassi {parentesi graffi} per inserire i merge tag.

Puoi anche includere allegati trascinandoli e rilasciandoli nell'email o utilizzando l'icona di allegato. I file provenienti dai campi personalizzati File possono essere automaticamente allegati se sono inferiori a 10MB.

Prima di finalizzare la tua automazione email, è consigliabile inviare un'email di prova a te stesso o a un collega per assicurarti che tutto funzioni come previsto.

## Casi d'Uso ed Esempi

Le automazioni email in Blue possono essere utilizzate per una varietà di scopi. Ecco alcuni esempi:

1. Invia un'email di conferma quando un cliente invia una richiesta tramite un modulo di raccolta. Imposta il trigger per inviare un'email quando un nuovo record viene creato tramite il modulo e assicurati di includere un campo email nel modulo per catturare l'indirizzo del cliente.
2. Notifica un assegnatario quando viene creata una nuova attività ad alta priorità. Imposta il trigger per inviare un'email quando un tag "Priorità" viene aggiunto a un record e utilizza il merge tag {Assignee} per inviare automaticamente l'email all'utente assegnato.
3. Invia un sondaggio a un cliente dopo che un ticket di supporto è stato contrassegnato come risolto. Imposta il trigger per inviare un'email quando un record è contrassegnato come completato e spostato nella lista "Fatto". Includi l'email del cliente in un campo personalizzato e fornisci informazioni dettagliate sul problema risolto nel corpo dell'email.
4. Automatizza un programma di reclutamento inviando email di conferma ai candidati. Imposta il trigger per inviare un'email quando un'applicazione viene inviata tramite un modulo e aggiunta alla lista "Ricevuta". Cattura l'email del candidato nel modulo e usala per inviare una risposta di ringraziamento.

## Vantaggi delle Automazioni Email

Le automazioni email in Blue offrono diversi vantaggi chiave:

- Comunicazione personalizzata attraverso l'uso di merge tag e dati dei campi personalizzati
- Notifiche automatiche che riducono il lavoro manuale e garantiscono aggiornamenti tempestivi
- Flussi di lavoro strutturati e basati sui dati che fanno avanzare i progetti in base ai dati dei record

## Conclusione

Le automazioni email in Blue sono uno strumento prezioso per semplificare la comunicazione e mantenere i progetti in carreggiata. Sfruttando i trigger, i merge tag e i dati dei campi personalizzati, puoi creare email automatizzate e personalizzate che migliorano la produttività del tuo team e garantiscono che gli aggiornamenti importanti non vengano mai persi. Con una vasta gamma di casi d'uso e una facile configurazione, le automazioni email sono una funzionalità indispensabile per qualsiasi utente di Blue che desideri ottimizzare il proprio flusso di lavoro.