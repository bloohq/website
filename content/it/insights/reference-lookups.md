---
title: Riferimenti e campi personalizzati di ricerca
description: Crea senza sforzo progetti interconnessi in Blue, trasformandolo in un'unica fonte di verità per la tua azienda con i nuovi campi di Riferimento e Ricerca.
category: "Product Updates"
date: 2023-11-01
---


I progetti in Blue sono già un modo potente per gestire i dati aziendali e far progredire il lavoro.

Oggi, stiamo facendo il passo logico successivo e permettendoti di interconnettere i tuoi dati *tra* i progetti per la massima flessibilità e potenza.

Interconnettere i progetti all'interno di Blue lo trasforma in un'unica fonte di verità per la tua azienda. Questa capacità consente la creazione di un dataset completo e interconnesso, abilitando un flusso di dati senza soluzione di continuità e una maggiore visibilità tra i progetti. Collegando i progetti, i team possono ottenere una visione unificata delle operazioni, migliorando il processo decisionale e l'efficienza operativa.

## Un esempio

Considera la ACME Company, che utilizza i campi personalizzati di Riferimento e Ricerca di Blue per creare un ecosistema di dati interconnesso tra i suoi progetti Clienti, Vendite e Inventario. I record dei clienti nel progetto Clienti sono collegati tramite campi di Riferimento alle transazioni di vendita nel progetto Vendite. Questo collegamento consente ai campi di Ricerca di estrarre i dettagli associati ai clienti, come numeri di telefono e stati degli account, direttamente in ciascun record di vendita. Inoltre, gli articoli di inventario venduti sono visualizzati nel record di vendita tramite un campo di Ricerca che fa riferimento ai dati della Quantità Venduta dal progetto Inventario. Infine, i prelievi di inventario sono collegati alle vendite pertinenti tramite un campo di Riferimento in Inventario, che punta ai record di Vendita. Questa configurazione fornisce piena visibilità su quale vendita ha attivato la rimozione dell'inventario, creando una vista integrata a 360 gradi tra i progetti.

## Come funzionano i campi di Riferimento

I campi personalizzati di Riferimento ti consentono di creare relazioni tra i record di progetti diversi in Blue. Quando crei un campo di Riferimento, l'Amministratore del Progetto seleziona il progetto specifico che fornirà l'elenco dei record di riferimento. Le opzioni di configurazione includono:

* **Selezione Singola**: Consente di scegliere un record di riferimento.
* **Selezione Multipla**: Consente di scegliere più record di riferimento.
* **Filtraggio**: Imposta filtri per consentire agli utenti di selezionare solo i record che corrispondono ai criteri di filtro.

Una volta configurato, gli utenti possono selezionare record specifici dal menu a discesa all'interno del campo di Riferimento, stabilendo un collegamento tra i progetti.

## Estendere i campi di riferimento utilizzando le ricerche

I campi personalizzati di Ricerca vengono utilizzati per importare dati da record in altri progetti, creando visibilità unidirezionale. Sono sempre di sola lettura e sono collegati a un campo personalizzato di Riferimento specifico. Quando un utente seleziona uno o più record utilizzando un campo personalizzato di Riferimento, il campo personalizzato di Ricerca mostrerà i dati di quei record. Le Ricerche possono visualizzare dati come:

* Creato il
* Aggiornato il
* Data di Scadenza
* Descrizione
* Elenco
* Etichetta
* Assegnatario
* Qualsiasi campo personalizzato supportato dal record di riferimento — inclusi altri campi di ricerca!

Ad esempio, immagina uno scenario in cui hai tre progetti: **Progetto A** è un progetto di vendite, **Progetto B** è un progetto di gestione dell'inventario e **Progetto C** è un progetto di gestione delle relazioni con i clienti. Nel Progetto A, hai un campo personalizzato di Riferimento che collega i record di vendita ai corrispondenti record dei clienti nel Progetto C. Nel Progetto B, hai un campo personalizzato di Ricerca che importa informazioni dal Progetto A, come la quantità venduta. In questo modo, quando viene creato un record di vendita nel Progetto A, le informazioni sui clienti associate a quella vendita vengono automaticamente estratte dal Progetto C, e la quantità venduta viene automaticamente estratta dal Progetto B. Questo ti consente di mantenere tutte le informazioni pertinenti in un unico posto e visualizzarle senza dover creare dati duplicati o aggiornare manualmente i record tra i progetti.

Un esempio reale di questo è un'azienda di e-commerce che utilizza Blue per gestire le proprie vendite, l'inventario e le relazioni con i clienti. Nel loro progetto **Vendite**, hanno un campo personalizzato di Riferimento che collega ciascun record di vendita al corrispondente record **Cliente** nel loro progetto **Clienti**. Nel loro progetto **Inventario**, hanno un campo personalizzato di Ricerca che importa informazioni dal progetto Vendite, come la quantità venduta, e le visualizza nel record dell'articolo di inventario. Questo consente loro di vedere facilmente quali vendite stanno guidando le rimozioni di inventario e mantenere i livelli di inventario aggiornati senza dover aggiornare manualmente i record tra i progetti.

## Conclusione

Immagina un mondo in cui i dati dei tuoi progetti non sono isolati, ma fluiscono liberamente tra i progetti, fornendo informazioni complete e aumentando l'efficienza. Questa è la potenza dei campi di Riferimento e Ricerca di Blue. Abilitando connessioni di dati senza soluzione di continuità e fornendo visibilità in tempo reale tra i progetti, queste funzionalità trasformano il modo in cui i team collaborano e prendono decisioni. Che tu stia gestendo relazioni con i clienti, tracciando vendite o supervisionando l'inventario, i campi di Riferimento e Ricerca in Blue consentono al tuo team di lavorare in modo più intelligente, veloce ed efficace. Immergiti nel mondo interconnesso di Blue e guarda la tua produttività decollare.

[Controlla la documentazione](https://documentation.blue.cc/custom-fields/reference) o [iscriviti e provalo tu stesso.](https://app.blue.cc)