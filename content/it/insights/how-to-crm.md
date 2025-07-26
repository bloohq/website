---
title: Come impostare Blue come CRM
description: Scopri come configurare Blue per tenere traccia dei tuoi clienti e delle tue trattative in modo semplice.
category: "Best Practices"
date: 2024-08-11
---


## Introduzione

Uno dei principali vantaggi dell'utilizzo di Blue è non usarlo per un *caso d'uso* specifico, ma usarlo *attraverso* casi d'uso. Questo significa che non devi pagare per più strumenti e hai anche un luogo dove puoi facilmente passare tra i tuoi vari progetti e processi come assunzioni, vendite, marketing e altro.

Nel corso degli anni, aiutando migliaia di clienti a configurarsi in Blue, abbiamo notato che la parte difficile non è *impostare* Blue stesso, ma riflettere sui processi e sfruttare al meglio la nostra piattaforma.

Le parti chiave sono pensare al flusso di lavoro passo dopo passo per ciascuno dei tuoi processi aziendali che desideri monitorare, e anche i dettagli dei dati che desideri catturare, e come questo si traduce nei campi personalizzati che imposti.

Oggi, ti guideremo nella creazione di [un sistema CRM di vendita facile da usare, ma potente](/solutions/use-case/sales-crm) con un database clienti collegato a un pipeline di opportunità. Tutti questi dati fluiranno in un dashboard dove puoi vedere dati in tempo reale sulle tue vendite totali, vendite previste e altro.

## Database Clienti

La prima cosa da fare è impostare un nuovo progetto per memorizzare i dati dei tuoi clienti. Questi dati verranno quindi incrociati in un altro progetto dove monitori opportunità di vendita specifiche.

Il motivo per cui separiamo le informazioni sui clienti dalle opportunità è che non si mappano uno a uno.

Un cliente può avere più opportunità o progetti.

Ad esempio, se sei un'agenzia di marketing e design, potresti inizialmente collaborare con un cliente per il suo branding, e poi fare un progetto separato per il suo sito web, e poi un altro per la gestione dei suoi social media.

Tutti questi sarebbero opportunità di vendita separate che richiedono il proprio monitoraggio e proposte, ma sono tutte collegate a quel singolo cliente.

Il vantaggio di separare il tuo database clienti in un progetto distinto è che se aggiorni qualsiasi dettaglio nel tuo database clienti, tutte le tue opportunità avranno automaticamente i nuovi dati, il che significa che ora hai una fonte di verità nella tua azienda! Non devi tornare indietro e modificare tutto manualmente!

Quindi, la prima cosa da decidere è se vuoi essere aziendale o incentrato sulla persona.

Questa decisione dipende davvero da cosa stai vendendo e a chi vendi. Se vendi principalmente a imprese, allora probabilmente vorrai che il nome del record sia il nome dell'azienda. Tuttavia, se vendi principalmente a individui (ad esempio, sei un coach di salute personale o un consulente di branding personale), allora probabilmente adotterai un approccio incentrato sulla persona.

Quindi il campo del nome del record sarà o il nome dell'azienda o il nome della persona, a seconda della tua scelta. Il motivo di questo è che significa che puoi facilmente identificare un cliente a colpo d'occhio, semplicemente guardando la tua bacheca o il tuo database.

Successivamente, devi considerare quali informazioni desideri catturare come parte del tuo database clienti. Questi diventeranno i tuoi campi personalizzati.

I soliti sospetti qui sono:

- Email
- Numero di telefono
- Sito web
- Indirizzo
- Fonte (ad es. da dove proviene questo cliente per la prima volta?)
- Categoria

In Blue, puoi anche rimuovere eventuali campi predefiniti di cui non hai bisogno. Per questo database clienti, ti consigliamo di rimuovere la data di scadenza, l'assegnatario, le dipendenze e le checklist. Potresti voler mantenere il nostro campo descrizione predefinito disponibile nel caso tu abbia note generali su quel cliente che non sono specifiche per alcuna opportunità di vendita.

Ti consigliamo di mantenere il campo "Riferito da", poiché sarà utile in seguito. Una volta impostato il nostro database delle opportunità, saremo in grado di vedere ogni record di vendita collegato a questo particolare cliente qui.

In termini di liste, di solito vediamo i nostri clienti mantenere tutto semplice e avere un'unica lista chiamata "Clienti" e lasciarla così. È meglio utilizzare tag o campi personalizzati per la categorizzazione.

Ciò che è fantastico qui è che una volta che hai questa configurazione, puoi facilmente importare i tuoi dati da altri sistemi o fogli Excel in Blue tramite la nostra funzione di importazione CSV, e puoi anche creare un modulo per i nuovi potenziali clienti per inviare i loro dettagli in modo da poterli **catturare automaticamente** nel tuo database.

## Database Opportunità

Ora che abbiamo il nostro database clienti, dobbiamo creare un altro progetto per catturare le nostre effettive opportunità di vendita. Puoi chiamare questo progetto "Sales CRM" o "Opportunità".

### Liste come Passaggi del Processo

Per impostare il tuo processo di vendita, devi pensare a quali sono i passaggi abituali che un'opportunità attraversa dal momento in cui ricevi una richiesta da un cliente fino a ottenere un contratto firmato.

Ogni lista nel tuo progetto sarà un passaggio nel tuo processo.

Indipendentemente dal tuo processo specifico, ci saranno alcune liste comuni che TUTTI i CRM di vendita dovrebbero avere:

- Non qualificato — Tutte le richieste in arrivo, dove non hai ancora qualificato un cliente.
- Chiuso Vinto — Tutte le opportunità che hai vinto e trasformato in vendite!
- Chiuso Perso — Tutte le opportunità in cui hai fatto un preventivo a un cliente e non hanno accettato.
- N/A — Qui metti tutte le opportunità che non hai vinto, ma che non sono state nemmeno "perse". Potrebbero essere quelle che hai rifiutato, quelle in cui il cliente, per qualsiasi motivo, ti ha ignorato, e così via.

In termini di riflessione sul tuo processo aziendale CRM di vendita, dovresti considerare il livello di granularità che desideri. Non ti consigliamo di avere 20 o 30 colonne, questo di solito diventa confuso e ti impedisce di vedere il quadro generale.

Tuttavia, è anche importante non rendere ogni processo troppo ampio, altrimenti le trattative si "bloccano" in una fase specifica per settimane o mesi, anche quando in realtà stanno avanzando. Ecco un approccio tipico raccomandato:

- **Non qualificato**: Tutte le richieste in arrivo, dove non hai ancora qualificato un cliente.
- **Qualificazione**: Qui prendi l'opportunità e inizi il processo di comprensione se questa è una buona corrispondenza per la tua azienda.
- **Scrittura della Proposta**: Qui inizi a trasformare l'opportunità in una proposta per la tua azienda. Questo è un documento che invieresti al cliente.
- **Proposta Inviata**: Qui hai inviato la proposta al cliente e stai aspettando una risposta.
- **Negoziazioni**: Qui sei nel processo di finalizzazione dell'affare.
- **Contratto in attesa di firma**: Qui stai solo aspettando che il cliente firmi il contratto.
- **Chiuso Vinto**: Qui hai vinto l'affare e ora stai lavorando al progetto.
- **Chiuso Perso**: Qui hai fatto un preventivo al cliente, ma non hanno accettato i termini.
- **N/A**: Qui metti tutte le opportunità che non hai vinto, ma che non sono state nemmeno "perse". Potrebbero essere quelle che hai rifiutato, quelle in cui il cliente, per qualsiasi motivo, ti ha ignorato, e così via.

### Tag come Categorie di Servizio
Parliamo ora dei tag.

Ti consigliamo di utilizzare tag per i diversi tipi di servizi che offri. Quindi, tornando al nostro esempio di agenzia di marketing e design, potresti avere tag per "branding", "sito web", "SEO", "Gestione Facebook", e così via.

I vantaggi qui sono che puoi facilmente filtrare per servizio con un clic, il che può darti una breve panoramica su quali servizi sono più popolari, e questo può anche informare future assunzioni, poiché tipicamente diversi servizi richiedono membri del team diversi.

### Campi Personalizzati del CRM di Vendita

Successivamente, dobbiamo considerare quali campi personalizzati vogliamo avere.

I campi tipici che vediamo utilizzati sono:

- **Importo**: Questo è un campo valuta per l'importo del progetto
- **Costo**: Il tuo costo previsto per soddisfare la vendita, anche un campo valuta
- **Profitto**: Un campo formula per calcolare il profitto basato sui campi importo e costo.
- **URL della Proposta**: Questo può includere un link a un documento Google Doc o Word online della tua proposta, in modo da poterlo facilmente cliccare e rivedere.
- **File ricevuti**: Questo può essere un campo file personalizzato dove puoi inserire eventuali file ricevuti dal cliente come materiali di ricerca, NDA, e così via.
- **Contratti**: Un altro campo file personalizzato dove puoi aggiungere contratti firmati per la custodia.
- **Livello di Fiducia**: Un campo personalizzato a stelle con 5 stelle, che indica quanto sei fiducioso di vincere questa particolare opportunità. Questo può essere utilizzato in seguito nel dashboard per le previsioni!
- **Data di Chiusura Prevista**: Un campo data per stimare quando l'affare è probabile che si chiuda.
- **Cliente**: Un campo di riferimento che collega alla persona di contatto principale nel database clienti.
- **Nome Cliente**: Un campo di ricerca che estrae il nome del cliente dal record collegato nel database clienti.
- **Email Cliente**: Un campo di ricerca che estrae l'email del cliente dal record collegato nel database clienti.
- **Fonte dell'Affare**: Un campo a discesa per monitorare da dove è originata l'opportunità (ad es., referral, sito web, chiamata a freddo, fiera).
- **Motivo Perso**: Un campo a discesa (per affari chiusi persi) per categorizzare il motivo per cui l'opportunità è stata persa.
- **Dimensione Cliente**: Un campo a discesa per categorizzare i clienti per dimensione (ad es., piccolo, medio, grande impresa).

Ancora una volta, è davvero **fino a te** decidere esattamente quali campi vuoi avere. Un avvertimento: è facile, durante la configurazione, aggiungere molti campi al tuo CRM di vendita per i dati che desideri catturare. Tuttavia, devi essere realistico in termini di disciplina e impegno di tempo. Non ha senso avere 30 campi nel tuo CRM di vendita se il 90% dei record non avrà dati in essi.

La cosa fantastica dei campi personalizzati è che si integrano bene in [Permessi Personalizzati](/platform/features/user-permissions). Questo significa che puoi decidere esattamente quali campi i membri del tuo team possono visualizzare o modificare. Ad esempio, potresti voler nascondere le informazioni sui costi e sui profitti dal personale junior.

### Automazioni

[Automazioni CRM di Vendita](/platform/features/automations) sono una funzionalità potente in Blue che può semplificare il tuo processo di vendita, garantire coerenza e risparmiare tempo su compiti ripetitivi. Impostando automazioni intelligenti, puoi migliorare l'efficacia del tuo CRM di vendita e consentire al tuo team di concentrarsi su ciò che conta di più: chiudere affari. Ecco alcune automazioni chiave da considerare per il tuo CRM di vendita:

- **Assegnazione Nuovo Lead**: Assegna automaticamente nuovi lead ai rappresentanti di vendita in base a criteri predefiniti come posizione, dimensione dell'affare o settore. Questo garantisce un rapido follow-up e una distribuzione equilibrata del carico di lavoro.
- **Promemoria di Follow-up**: Imposta promemoria automatizzati per i rappresentanti di vendita per seguire i potenziali clienti dopo un certo periodo di inattività. Questo aiuta a prevenire che i lead vengano trascurati.
- **Notifiche di Progresso Fase**: Notifica ai membri del team pertinenti quando un affare passa a una nuova fase nel pipeline. Questo tiene tutti informati sui progressi e consente interventi tempestivi se necessario.
- **Avvisi di Invecchiamento Affari**: Crea avvisi per affari che sono stati in una particolare fase più a lungo del previsto. Questo aiuta a identificare affari bloccati che potrebbero necessitare di ulteriore attenzione.

## Collegare Clienti e Affari

Una delle funzionalità più potenti di Blue per creare un sistema CRM efficace è la possibilità di collegare il tuo database clienti con le tue opportunità di vendita. Questa connessione ti consente di mantenere una singola fonte di verità per le informazioni sui clienti mentre monitori più affari associati a ciascun cliente. Esploriamo come impostare questo utilizzando campi personalizzati di Riferimento e Ricerca.

### Impostare il Campo di Riferimento

1. Nel tuo progetto Opportunità (o Sales CRM), crea un nuovo campo personalizzato.
2. Scegli il tipo di campo "Riferimento".
3. Seleziona il tuo progetto Database Clienti come fonte per il riferimento.
4. Configura il campo per consentire la selezione singola (poiché ogni opportunità è tipicamente associata a un cliente).
5. Nomina questo campo qualcosa come "Cliente" o "Azienda Associata".

Ora, quando crei o modifichi un'opportunità, sarai in grado di selezionare il cliente associato da un menu a discesa popolato con record dal tuo Database Clienti.

### Migliorare con i Campi di Ricerca

Una volta stabilita la connessione di riferimento, puoi utilizzare i campi di Ricerca per portare informazioni rilevanti sui clienti direttamente nella tua vista delle opportunità. Ecco come:

1. Nel tuo progetto Opportunità, crea un nuovo campo personalizzato.
2. Scegli il tipo di campo "Ricerca".
3. Seleziona il campo di Riferimento che hai appena creato ("Cliente" o "Azienda Associata") come fonte.
4. Scegli quali informazioni sui clienti desideri visualizzare. Potresti considerare campi come: Email, Numero di Telefono, Categoria Cliente, Manager Account.

Ripeti questo processo per ciascun pezzo di informazione sui clienti che desideri visualizzare nella tua vista delle opportunità.

I vantaggi di questo sono:

- **Singola Fonte di Verità**: Aggiorna le informazioni sui clienti una volta nel Database Clienti, e si riflettono automaticamente in tutte le opportunità collegate.
- **Efficienza**: Accedi rapidamente ai dettagli rilevanti sui clienti mentre lavori sulle opportunità senza dover passare tra progetti.
- **Integrità dei Dati**: Riduci gli errori da inserimento manuale dei dati estraendo automaticamente le informazioni sui clienti.
- **Vista Olistica**: Vedi facilmente tutte le opportunità associate a un cliente utilizzando il campo "Riferito da" nel tuo Database Clienti.

### Suggerimento Avanzato: Ricerca di una Ricerca

Blue offre una funzionalità avanzata chiamata "Ricerca di una Ricerca" che può essere incredibilmente utile per configurazioni CRM complesse. Questa funzionalità ti consente di creare connessioni attraverso più progetti, consentendoti di accedere alle informazioni sia dal tuo Database Clienti che dal progetto Opportunità in un terzo progetto.

Ad esempio, supponiamo di avere uno spazio di lavoro "Progetti" dove gestisci il lavoro effettivo per i tuoi clienti. Vuoi che questo spazio di lavoro abbia accesso sia ai dettagli dei clienti che alle informazioni sulle opportunità. Ecco come puoi impostarlo:

Per prima cosa, crea un campo di Riferimento nel tuo spazio di lavoro Progetti che si collega al progetto Opportunità. Questo stabilisce la connessione iniziale. Successivamente, crea campi di Ricerca basati su questo Riferimento per estrarre dettagli specifici dalle opportunità, come il valore dell'affare o la data di chiusura prevista.

Il vero potere arriva nel passaggio successivo: puoi creare ulteriori campi di Ricerca che raggiungono attraverso il Riferimento dell'opportunità al Database Clienti. Questo ti consente di estrarre informazioni sui clienti come dettagli di contatto o stato dell'account direttamente nel tuo spazio di lavoro Progetti.

Questa catena di connessioni ti offre una vista completa nel tuo spazio di lavoro Progetti, combinando dati sia dalle tue opportunità che dal database clienti. È un modo potente per garantire che i tuoi team di progetto abbiano tutte le informazioni rilevanti a portata di mano senza dover passare tra diversi progetti.

### Migliori Pratiche per Sistemi CRM Collegati

Mantieni il tuo Database Clienti come la singola fonte di verità per tutte le informazioni sui clienti. Ogni volta che hai bisogno di aggiornare i dettagli sui clienti, fallo sempre prima nel Database Clienti. Questo garantisce che le informazioni rimangano coerenti in tutti i progetti collegati.

Quando crei campi di Riferimento e Ricerca, utilizza nomi chiari e significativi. Questo aiuta a mantenere la chiarezza, specialmente man mano che il tuo sistema cresce in complessità.

Rivedi regolarmente la tua configurazione per assicurarti di estrarre le informazioni più rilevanti. Man mano che le esigenze della tua azienda evolvono, potresti dover aggiungere nuovi campi di Ricerca o rimuovere quelli che non sono più utili. Le revisioni periodiche aiutano a mantenere il tuo sistema snello ed efficace.

Considera di sfruttare le funzionalità di automazione di Blue per mantenere i tuoi dati sincronizzati e aggiornati tra i progetti. Ad esempio, potresti impostare un'automazione per notificare i membri del team pertinenti quando le informazioni chiave sui clienti vengono aggiornate nel Database Clienti.

Implementando efficacemente queste strategie e sfruttando appieno i campi di Riferimento e Ricerca, puoi creare un potente sistema CRM interconnesso in Blue. Questo sistema ti fornirà una visione completa a 360 gradi delle tue relazioni con i clienti e del tuo pipeline di vendita, consentendo decisioni più informate e operazioni più fluide in tutta la tua organizzazione.

## Dashboard

Le dashboard sono un componente cruciale di qualsiasi sistema CRM efficace, fornendo informazioni immediate sulle tue performance di vendita e relazioni con i clienti. La funzionalità dashboard di Blue è particolarmente potente perché ti consente di combinare dati in tempo reale provenienti da più progetti contemporaneamente, offrendoti una visione completa delle tue operazioni di vendita.

Quando imposti il tuo dashboard CRM in Blue, considera di includere diversi metriche chiave. Il pipeline generato per mese mostra il valore totale delle nuove opportunità aggiunte al tuo pipeline, aiutandoti a monitorare la capacità del tuo team di generare nuovo business. Le vendite per mese mostrano i tuoi affari effettivamente chiusi, consentendoti di monitorare le performance del tuo team nel convertire opportunità in vendite.

Introdurre il concetto di sconti sul pipeline può portare a previsioni più accurate. Ad esempio, potresti contare il 90% del valore degli affari nella fase "Contratto in attesa di firma", ma solo il 50% degli affari nella fase "Proposta inviata". Questo approccio ponderato fornisce una previsione di vendita più realistica.

Monitorare le nuove opportunità per mese ti aiuta a tenere traccia del numero di nuovi potenziali affari che entrano nel tuo pipeline, il che è un buon indicatore degli sforzi di prospecting del tuo team di vendita. Suddividere le vendite per tipo può aiutarti a identificare le tue offerte più riuscite. Se imposti un progetto di monitoraggio delle fatture collegato alle tue opportunità, puoi anche monitorare il fatturato effettivo nel tuo dashboard, fornendo un quadro completo dall'opportunità al cash.

Blue offre diverse funzionalità potenti per aiutarti a creare un dashboard CRM informativo e interattivo. La piattaforma fornisce tre principali tipi di grafici: schede statistiche, grafici a torta e grafici a barre. Le schede statistiche sono ideali per visualizzare metriche chiave come il valore totale del pipeline o il numero di opportunità attive. I grafici a torta sono perfetti per mostrare la composizione delle tue vendite per tipo o la distribuzione degli affari attraverso le diverse fasi. I grafici a barre eccellono nel confrontare metriche nel tempo, come vendite mensili o nuove opportunità.

Le sofisticate capacità di filtraggio di Blue ti consentono di segmentare i tuoi dati per progetto, lista, tag e intervallo di tempo. Questo è particolarmente utile per approfondire aspetti specifici dei tuoi dati di vendita o confrontare le performance tra diversi team o prodotti. La piattaforma consolida intelligentemente liste e tag con lo stesso nome tra i progetti, consentendo un'analisi senza soluzione di continuità tra i progetti. Questo è inestimabile per una configurazione CRM in cui potresti avere progetti separati per clienti, opportunità e fatture.

La personalizzazione è una forza chiave della funzionalità dashboard di Blue. La funzionalità di trascinamento e rilascio e la flessibilità di visualizzazione ti consentono di creare un dashboard che si adatta perfettamente alle tue esigenze. Puoi facilmente riordinare i grafici e scegliere la visualizzazione più appropriata per ciascuna metrica.
Sebbene le dashboard siano attualmente utilizzate solo internamente, puoi facilmente condividerle con i membri del team, concedendo permessi di visualizzazione o modifica. Questo garantisce che tutti nel tuo team di vendita abbiano accesso alle informazioni di cui hanno bisogno.

Sfruttando queste funzionalità e includendo le metriche chiave di cui abbiamo discusso, puoi creare un dashboard CRM completo in Blue che fornisce informazioni in tempo reale sulle tue performance di vendita, salute del pipeline e crescita complessiva del business. Questo dashboard diventerà uno strumento prezioso per prendere decisioni basate sui dati e mantenere l'intero team allineato sui tuoi obiettivi e progressi di vendita.

## Conclusione

Impostare un CRM di vendita completo in Blue è un modo potente per semplificare il tuo processo di vendita e ottenere preziose informazioni sulle tue relazioni con i clienti e sulle performance aziendali. Seguendo i passaggi delineati in questa guida, hai creato un sistema robusto che integra informazioni sui clienti, opportunità di vendita e metriche di performance in un'unica piattaforma coesa.

Abbiamo iniziato creando un database clienti, stabilendo una singola fonte di verità per tutte le informazioni sui tuoi clienti. Questa base ti consente di mantenere registri accurati e aggiornati per tutti i tuoi clienti e potenziali. Abbiamo poi costruito su questo con un database delle opportunità, consentendoti di monitorare e gestire efficacemente il tuo pipeline di vendita.

Uno dei punti di forza chiave dell'utilizzo di Blue per il tuo CRM è la possibilità di collegare questi database utilizzando campi di riferimento e ricerca. Questa integrazione crea un sistema dinamico in cui gli aggiornamenti alle informazioni sui clienti si riflettono istantaneamente in tutte le opportunità correlate, garantendo coerenza dei dati e risparmiando tempo sugli aggiornamenti manuali.
Abbiamo esplorato come sfruttare le potenti funzionalità di automazione di Blue per semplificare il tuo flusso di lavoro, dall'assegnazione di nuovi lead all'invio di promemoria di follow-up. Queste automazioni aiutano a garantire che nessuna opportunità venga trascurata e che il tuo team possa concentrarsi su attività ad alto valore piuttosto che su compiti amministrativi.

Infine, abbiamo approfondito la creazione di dashboard che forniscono informazioni immediate sulle tue performance di vendita. Combinando dati dai tuoi database clienti e opportunità, queste dashboard offrono una visione completa del tuo pipeline di vendita, affari chiusi e salute complessiva del business.

Ricorda, la chiave per ottenere il massimo dal tuo CRM è un uso costante e un affinamento regolare. Incoraggia il tuo team ad adottare completamente il sistema, rivedi regolarmente i tuoi processi e automazioni e continua a esplorare nuovi modi per sfruttare le funzionalità di Blue a supporto dei tuoi sforzi di vendita.

Con questa configurazione del CRM di vendita in Blue, sei ben attrezzato per coltivare relazioni con i clienti, chiudere più affari e far progredire la tua azienda.