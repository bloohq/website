---
title: Creare checklist riutilizzabili utilizzando automazioni
description: Scopri come creare automazioni di gestione dei progetti per checklist riutilizzabili.
category: "Best Practices"
date: 2024-07-08
---


In molti progetti e processi, potrebbe essere necessario utilizzare la stessa checklist su più record o attività.

Tuttavia, non è molto efficiente riscrivere manualmente la checklist ogni volta che si desidera aggiungerla a un record. È qui che puoi sfruttare [potenti automazioni di gestione dei progetti](/platform/features/automations) per farlo automaticamente per te!

Ricorda che le automazioni in Blue richiedono due cose fondamentali:

1. Un Trigger — Cosa dovrebbe accadere per avviare l'automazione. Questo può essere quando a un record viene assegnato un tag specifico, quando si sposta in un elenco specifico.
2. Una o più Azioni — In questo caso, sarebbe la creazione automatica di una o più checklist.

Iniziamo con l'azione prima, poi discuteremo dei possibili trigger che puoi utilizzare.

## Azione di Automazione della Checklist

Puoi creare una nuova automazione e impostare una o più checklist da creare, come nell'esempio qui sotto:

![](/insights/checklist-automation.png)

Queste sarebbero le checklist che desideri vengano create ogni volta che esegui l'azione.

## Trigger di Automazione della Checklist

Ci sono diversi modi per attivare la creazione delle tue checklist riutilizzabili. Ecco alcune opzioni popolari:

- **Aggiunta di un Tag Specifico:** Puoi impostare l'automazione per attivarsi quando viene aggiunto un tag particolare a un record. Ad esempio, quando viene aggiunto il tag "Nuovo Progetto", potrebbe automaticamente creare la tua checklist di avvio del progetto.
- **Assegnazione del Record:** La creazione della checklist può essere attivata quando un record viene assegnato a un individuo specifico o a chiunque. Questo è utile per checklist di onboarding o procedure specifiche per attività.
- **Spostamento in un Elenco Specifico:** Quando un record viene spostato in un elenco particolare nel tuo progetto, può attivare la creazione di una checklist pertinente. Ad esempio, spostare un elemento in un elenco "Controllo Qualità" potrebbe attivare una checklist di QA.
- **Campo Checkbox Personalizzato:** Crea un campo checkbox personalizzato e imposta l'automazione per attivarsi quando questa casella viene selezionata. Questo ti dà il controllo manuale su quando aggiungere la checklist.
- **Campi Personalizzati a Selezione Singola o Multipla:** Puoi creare un campo personalizzato a selezione singola o multipla con varie opzioni. Ogni opzione può essere collegata a un modello di checklist specifico tramite automazioni separate. Questo consente un controllo più granulare e la possibilità di avere più modelli di checklist pronti per diversi scenari.

Per migliorare il controllo su chi può attivare queste automazioni, puoi nascondere questi campi personalizzati da determinati utenti utilizzando ruoli utente personalizzati. Questo garantisce che solo gli amministratori di progetto o altro personale autorizzato possano attivare queste opzioni.

Ricorda, la chiave per un uso efficace delle checklist riutilizzabili con automazioni è progettare i tuoi trigger in modo ponderato. Considera il flusso di lavoro del tuo team, i tipi di progetti che gestisci e chi dovrebbe avere la possibilità di avviare diversi processi. Con automazioni ben pianificate, puoi semplificare notevolmente la gestione dei tuoi progetti e garantire coerenza nelle tue operazioni.

## Risorse Utili

- [Documentazione sull'Automazione della Gestione dei Progetti](https://documentation.blue.cc/automations)
- [Documentazione sui Ruoli Utente Personalizzati](https://documentation.blue.cc/user-management/roles/custom-user-roles)
- [Documentazione sui Campi Personalizzati](https://documentation.blue.cc/custom-fields)