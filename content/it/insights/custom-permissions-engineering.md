---
title: Creare il motore di autorizzazioni personalizzate di Blue
description: Scopri cosa c'è dietro le quinte con il team di ingegneria di Blue mentre spiegano come costruire una funzionalità di categorizzazione automatica e tagging alimentata dall'IA.
category: "Engineering"
date: 2024-07-25
---


La gestione efficace dei progetti e dei processi è cruciale per le organizzazioni di tutte le dimensioni.

In Blue, [abbiamo fatto della nostra missione](/about) organizzare il lavoro nel mondo costruendo la migliore piattaforma di gestione dei progetti del pianeta: semplice, potente, flessibile e accessibile a tutti.

Ciò significa che la nostra piattaforma deve adattarsi alle esigenze uniche di ciascun team. Oggi siamo entusiasti di sollevare il velo su una delle nostre funzionalità più potenti: le autorizzazioni personalizzate.

Gli strumenti di gestione dei progetti sono la spina dorsale dei flussi di lavoro moderni, ospitando dati sensibili, comunicazioni cruciali e piani strategici. Pertanto, la capacità di controllare finemente l'accesso a queste informazioni non è solo un lusso, ma una necessità.

<video autoplay loop muted playsinline>
  <source src="/videos/user-roles.mp4" type="video/mp4">
</video>

Le autorizzazioni personalizzate svolgono un ruolo fondamentale nelle piattaforme B2B SaaS, specialmente negli strumenti di gestione dei progetti, dove l'equilibrio tra collaborazione e sicurezza può determinare il successo di un progetto.

Ma qui Blue adotta un approccio diverso: **crediamo che le funzionalità di livello enterprise non debbano essere riservate a budget di grandi dimensioni.**

In un'era in cui l'IA consente a piccoli team di operare su scale senza precedenti, perché la sicurezza robusta e la personalizzazione dovrebbero essere fuori portata?

In questo sguardo dietro le quinte, esploreremo come abbiamo sviluppato la nostra funzionalità di autorizzazioni personalizzate, sfidando lo status quo dei livelli di prezzo SaaS e portando opzioni di sicurezza potenti e flessibili alle aziende di tutte le dimensioni.

Che tu sia una startup con grandi sogni o un attore consolidato che cerca di ottimizzare i propri processi, le autorizzazioni personalizzate possono abilitare nuovi casi d'uso che non sapevi nemmeno fossero possibili.

## Comprendere le autorizzazioni utente personalizzate

Prima di immergerci nel nostro viaggio di sviluppo delle autorizzazioni personalizzate per Blue, prendiamoci un momento per capire cosa sono le autorizzazioni utente personalizzate e perché sono così cruciali nel software di gestione dei progetti.

Le autorizzazioni utente personalizzate si riferiscono alla capacità di personalizzare i diritti di accesso per singoli utenti o gruppi all'interno di un sistema software. Invece di fare affidamento su ruoli predefiniti con set di autorizzazioni fissi, le autorizzazioni personalizzate consentono agli amministratori di creare profili di accesso altamente specifici che si allineano perfettamente con la struttura e le esigenze di flusso di lavoro della loro organizzazione.

Nel contesto di software di gestione dei progetti come Blue, le autorizzazioni personalizzate includono:

* **Controllo degli accessi granulare**: Determinare chi può visualizzare, modificare o eliminare specifici tipi di dati di progetto.
* **Restrizioni basate sulle funzionalità**: Abilitare o disabilitare determinate funzionalità per utenti o team specifici.
* **Livelli di sensibilità dei dati**: Impostare diversi livelli di accesso a informazioni sensibili all'interno dei progetti.
* **Autorizzazioni specifiche per flussi di lavoro**: Allineare le capacità degli utenti con fasi o aspetti specifici del flusso di lavoro del progetto.

L'importanza delle autorizzazioni personalizzate nella gestione dei progetti non può essere sottovalutata:

* **Sicurezza migliorata**: Fornendo agli utenti solo l'accesso di cui hanno bisogno, si riduce il rischio di violazioni dei dati o modifiche non autorizzate.
* **Conformità migliorata**: Le autorizzazioni personalizzate aiutano le organizzazioni a soddisfare i requisiti normativi specifici del settore controllando l'accesso ai dati.
* **Collaborazione semplificata**: I team possono lavorare in modo più efficiente quando ogni membro ha il giusto livello di accesso per svolgere il proprio ruolo senza restrizioni non necessarie o privilegi eccessivi.
* **Flessibilità per organizzazioni complesse**: Man mano che le aziende crescono ed evolvono, le autorizzazioni personalizzate consentono al software di adattarsi a strutture e processi organizzativi in cambiamento.

## Arrivare a un SÌ

[Abbiamo scritto in precedenza](/insights/value-proposition-blue) che ogni funzionalità in Blue deve essere un **SÌ** deciso prima di decidere di costruirla. Non abbiamo il lusso di avere centinaia di ingegneri e di sprecare tempo e denaro costruendo cose di cui i clienti non hanno bisogno.

E così, il percorso per implementare le autorizzazioni personalizzate in Blue non è stato una linea retta. Come molte funzionalità potenti, è iniziato con una chiara esigenza da parte dei nostri utenti ed è evoluto attraverso una considerazione e una pianificazione attente.

Per anni, i nostri clienti avevano richiesto un controllo più granulare sulle autorizzazioni degli utenti. Man mano che le organizzazioni di tutte le dimensioni iniziavano a gestire progetti sempre più complessi e sensibili, le limitazioni del nostro standard di controllo degli accessi basato sui ruoli diventavano evidenti.

Piccole startup che lavorano con clienti esterni, aziende di medie dimensioni con processi di approvazione intricati e grandi imprese con rigorosi requisiti di conformità hanno tutti espresso la stessa necessità:

Maggiore flessibilità nella gestione dell'accesso degli utenti.

Nonostante la chiara domanda, inizialmente abbiamo esitato a immergerci nello sviluppo delle autorizzazioni personalizzate.

Perché?

Comprendevamo la complessità coinvolta!

Le autorizzazioni personalizzate toccano ogni parte di un sistema di gestione dei progetti, dall'interfaccia utente fino alla struttura del database. Sapevamo che implementare questa funzionalità avrebbe richiesto cambiamenti significativi alla nostra architettura di base e una considerazione attenta delle implicazioni sulle prestazioni.

Mentre esaminavamo il panorama, notammo che molto pochi dei nostri concorrenti avevano tentato di implementare un potente motore di autorizzazioni personalizzate come quello richiesto dai nostri clienti. Coloro che lo facevano spesso lo riservavano ai loro piani enterprise di livello più alto.

Diventò chiaro il perché: lo sforzo di sviluppo è sostanziale e le poste in gioco sono elevate.

Implementare le autorizzazioni personalizzate in modo errato potrebbe introdurre bug critici o vulnerabilità di sicurezza, compromettendo potenzialmente l'intero sistema. Questa realizzazione sottolineò l'entità della sfida che stavamo considerando.

### Sfida allo Status Quo

Tuttavia, mentre continuavamo a crescere ed evolverci, raggiungemmo una realizzazione fondamentale:

**Il modello SaaS tradizionale di riservare funzionalità potenti ai clienti enterprise non ha più senso nel panorama aziendale odierno.**

Nel 2024, con il potere dell'IA e strumenti avanzati, i piccoli team possono operare su una scala e complessità che rivaleggiano con organizzazioni molto più grandi. Una startup potrebbe gestire dati sensibili dei clienti in più paesi. Un'agenzia di marketing di piccole dimensioni potrebbe gestire dozzine di progetti per clienti con requisiti di riservatezza variabili. Queste aziende hanno bisogno dello stesso livello di sicurezza e personalizzazione di *qualsiasi* grande impresa.

Ci siamo chiesti: perché la dimensione della forza lavoro o del budget di un'azienda dovrebbe determinare la loro capacità di mantenere i propri dati al sicuro e i propri processi efficienti?

### Funzionalità di livello enterprise per tutti

Questa realizzazione ci ha portato a una filosofia fondamentale che ora guida gran parte del nostro sviluppo in Blue: le funzionalità di livello enterprise dovrebbero essere accessibili alle aziende di tutte le dimensioni.

Crediamo che:

- **La sicurezza non dovrebbe essere un lusso.** Ogni azienda, indipendentemente dalle dimensioni, merita gli strumenti per proteggere i propri dati e processi.
- **La flessibilità guida l'innovazione.** Dando a tutti i nostri utenti strumenti potenti, consentiamo loro di creare flussi di lavoro e sistemi che spingono le loro industrie in avanti.
- **La crescita non dovrebbe richiedere cambiamenti nella piattaforma.** Man mano che i nostri clienti crescono, i loro strumenti dovrebbero crescere senza soluzione di continuità con loro.

Con questa mentalità, abbiamo deciso di affrontare la sfida delle autorizzazioni personalizzate a viso aperto, impegnandoci a renderle disponibili a tutti i nostri utenti, non solo a quelli con piani di livello superiore.

Questa decisione ci ha messo su un percorso di progettazione attenta, sviluppo iterativo e feedback continuo degli utenti che alla fine ha portato alla funzionalità di autorizzazioni personalizzate di cui siamo orgogliosi di offrire oggi.

Nella sezione successiva, ci immergeremo in come abbiamo affrontato il processo di progettazione e sviluppo per dare vita a questa funzionalità complessa.

### Progettazione e Sviluppo

Quando abbiamo deciso di affrontare le autorizzazioni personalizzate, ci siamo rapidamente resi conto di trovarci di fronte a un compito colossale.

A prima vista, "autorizzazioni personalizzate" potrebbe sembrare semplice, ma è una funzionalità ingannevolmente complessa che tocca ogni aspetto del nostro sistema.

La sfida era scoraggiante: dovevamo implementare autorizzazioni a cascata, consentire modifiche al volo, apportare cambiamenti significativi allo schema del database e garantire un funzionamento senza soluzione di continuità in tutto il nostro ecosistema: web, Mac, Windows, app iOS e Android, così come la nostra API e i webhook.

La complessità era sufficiente a far fermare anche i programmatori più esperti.

Il nostro approccio si è incentrato su due principi chiave:

1. Suddividere la funzionalità in versioni gestibili
2. Abbracciare la spedizione incrementale.

Di fronte alla complessità delle autorizzazioni personalizzate su larga scala, ci siamo posti una domanda cruciale:

> Quale sarebbe la versione più semplice possibile di questa funzionalità?

Questo approccio si allinea con il principio agile di fornire un Minimum Viable Product (MVP) e iterare in base al feedback.

La nostra risposta è stata sorprendentemente semplice:

1. Introdurre un interruttore per nascondere la scheda di attività del progetto
2. Aggiungere un altro interruttore per nascondere la scheda dei moduli

**Ecco tutto.**

Nessun campanello e fischietto, nessuna matrice di autorizzazioni complessa: solo due semplici interruttori on/off.

Sebbene possa sembrare deludente a prima vista, questo approccio ha offerto diversi vantaggi significativi:

* **Implementazione rapida**: Questi semplici interruttori potrebbero essere sviluppati e testati rapidamente, consentendoci di mettere una versione base delle autorizzazioni personalizzate nelle mani degli utenti in tempi brevi.
* **Valore chiaro per l'utente**: Anche con solo queste due opzioni, stavamo fornendo un valore tangibile. Alcuni team potrebbero voler nascondere il feed di attività dai clienti, mentre altri potrebbero aver bisogno di limitare l'accesso ai moduli per determinati gruppi di utenti.
* **Fondamento per la crescita**: Questo semplice inizio ha gettato le basi per autorizzazioni più complesse. Ci ha permesso di impostare l'infrastruttura di base per le autorizzazioni personalizzate senza essere sopraffatti dalla complessità fin dall'inizio.
* **Feedback degli utenti**: Rilasciando questa versione semplice, potevamo raccogliere feedback reali su come gli utenti interagivano con le autorizzazioni personalizzate, informando il nostro sviluppo futuro.
* **Apprendimento tecnico**: Questa implementazione iniziale ha fornito al nostro team di sviluppo esperienza pratica nella modifica delle autorizzazioni attraverso la nostra piattaforma, preparandoci per iterazioni più complesse.

E sai, è davvero umiliante avere una grande visione per qualcosa e poi spedire qualcosa che è una percentuale così piccola di quella visione.

Dopo aver spedito questi primi due interruttori, abbiamo deciso di affrontare qualcosa di più sofisticato. Siamo arrivati a due nuove autorizzazioni per ruoli utente personalizzati.

La prima era la possibilità di limitare gli utenti a visualizzare solo i record che sono stati specificamente assegnati a loro. Questo è molto utile se hai un cliente in un progetto e vuoi che veda solo i record specificamente assegnati a lui invece di tutto ciò su cui stai lavorando per lui.

La seconda era un'opzione per gli amministratori di progetto per bloccare i gruppi di utenti dall'invitare altri utenti. Questo è utile se hai un progetto sensibile che vuoi garantire rimanga su una base di "necessità di vedere".

Una volta che abbiamo spedito questo, abbiamo guadagnato maggiore fiducia e per la nostra terza versione abbiamo affrontato le autorizzazioni a livello di colonna, il che significa essere in grado di decidere quali campi personalizzati un gruppo di utenti specifico può visualizzare o modificare.

Questo è estremamente potente. Immagina di avere un progetto CRM e di avere dati che non sono solo relativi agli importi che il cliente pagherà, ma anche ai tuoi costi e margini di profitto. Potresti non voler che i tuoi campi di costo e il campo della formula del margine di progetto siano visibili al personale junior, e le autorizzazioni personalizzate ti consentono di bloccare quei campi in modo che non vengano mostrati.

Successivamente, ci siamo spostati sulla creazione di autorizzazioni basate su liste, dove gli amministratori di progetto possono decidere se un gruppo di utenti può visualizzare, modificare ed eliminare una lista specifica. Se nascondono una lista, tutti i record all'interno di quella lista diventano anch'essi nascosti, il che è fantastico perché significa che puoi nascondere determinate parti del tuo processo dai membri del team o dai clienti.

Questo è il risultato finale:

<video autoplay loop muted playsinline>
  <source src="/videos/custom-user-roles.mp4" type="video/mp4">
</video>

## Considerazioni Tecniche

Al centro dell'architettura tecnica di Blue c'è GraphQL, una scelta fondamentale che ha influenzato significativamente la nostra capacità di implementare funzionalità complesse come le autorizzazioni personalizzate. Ma prima di immergerci nei dettagli, facciamo un passo indietro e comprendiamo cos'è GraphQL e come si differenzia dall'approccio più tradizionale delle API REST.
GraphQL vs API REST: Una spiegazione accessibile

Immagina di essere in un ristorante. Con un'API REST, è come ordinare da un menu fisso. Chiedi un piatto specifico (endpoint) e ottieni tutto ciò che lo accompagna, che tu lo voglia o meno. Se desideri personalizzare il tuo pasto, potresti dover fare più ordini (chiamate API) o chiedere un piatto preparato appositamente (endpoint personalizzato).

GraphQL, d'altra parte, è come avere una conversazione con uno chef che può preparare qualsiasi cosa. Dici allo chef esattamente quali ingredienti desideri (campi dati) e in quali quantità. Lo chef poi prepara un piatto che è esattamente ciò che hai chiesto: né di più né di meno. Questo è essenzialmente ciò che fa GraphQL: consente al client di richiedere esattamente i dati di cui ha bisogno e il server fornisce solo quelli.

### Un Pranzo Importante

Circa sei settimane dopo l'inizio dello sviluppo di Blue, il nostro ingegnere capo e CEO sono andati a pranzo.

L'argomento di discussione?

Se passare dalle API REST a GraphQL. Questa non era una decisione da prendere alla leggera: adottare GraphQL avrebbe significato scartare sei settimane di lavoro iniziale.

Durante il ritorno in ufficio, il CEO ha posto una domanda cruciale all'ingegnere capo: "Ci pentiremo di non averlo fatto tra cinque anni?"

La risposta è diventata chiara: GraphQL era la strada da seguire.

Abbiamo riconosciuto il potenziale di questa tecnologia fin dall'inizio, vedendo come potesse supportare la nostra visione per una piattaforma di gestione dei progetti flessibile e potente.

La nostra lungimiranza nell'adottare GraphQL ha dato i suoi frutti quando si è trattato di implementare le autorizzazioni personalizzate. Con un'API REST, avremmo avuto bisogno di un endpoint diverso per ogni possibile configurazione di autorizzazioni personalizzate, un approccio che sarebbe rapidamente diventato ingombrante e difficile da mantenere.

GraphQL, tuttavia, ci consente di gestire le autorizzazioni personalizzate in modo dinamico. Ecco come funziona:

- **Controlli delle autorizzazioni al volo**: Quando un client effettua una richiesta, il nostro server GraphQL può controllare le autorizzazioni dell'utente direttamente dal nostro database.
- **Recupero dati preciso**: In base a queste autorizzazioni, GraphQL restituisce solo i dati richiesti che rientrano nei diritti di accesso dell'utente.
- **Query flessibili**: Man mano che le autorizzazioni cambiano, non è necessario creare nuovi endpoint o modificare quelli esistenti. La stessa query GraphQL può adattarsi a diverse configurazioni di autorizzazioni.
- **Recupero dati efficiente**: GraphQL consente ai client di richiedere esattamente ciò di cui hanno bisogno. Ciò significa che non stiamo sovra-recuperando dati, il che potrebbe esporre informazioni a cui l'utente non dovrebbe avere accesso.

Questa flessibilità è cruciale per una funzionalità complessa come le autorizzazioni personalizzate. Ci consente di offrire un controllo granulare *senza* sacrificare le prestazioni o la manutenibilità.

## Sfide

Implementare le autorizzazioni personalizzate in Blue ha portato con sé una serie di sfide, ognuna delle quali ci ha spinto a innovare e affinare il nostro approccio. L'ottimizzazione delle prestazioni è rapidamente emersa come una preoccupazione critica. Man mano che aggiungevamo controlli delle autorizzazioni più granulari, rischiavamo di rallentare il nostro sistema, specialmente per progetti di grandi dimensioni con molti utenti e configurazioni di autorizzazioni complesse. Per affrontare questo, abbiamo implementato una strategia di caching multi-livello, ottimizzato le nostre query di database e sfruttato la capacità di GraphQL di richiedere solo i dati necessari. Questo approccio ci ha permesso di mantenere tempi di risposta rapidi anche man mano che i progetti si espandevano e la complessità delle autorizzazioni cresceva.

L'interfaccia utente per le autorizzazioni personalizzate ha presentato un'altra significativa difficoltà. Dovevamo rendere l'interfaccia intuitiva e gestibile per gli amministratori, anche mentre aggiungevamo più opzioni e aumentavamo la complessità del sistema.

La nostra soluzione ha coinvolto più cicli di test degli utenti e design iterativo.

Abbiamo introdotto una matrice visiva delle autorizzazioni che consentiva agli amministratori di visualizzare e modificare rapidamente le autorizzazioni attraverso diversi ruoli e aree di progetto.

Garantire la coerenza tra le piattaforme ha presentato un proprio insieme di sfide. Dovevamo implementare le autorizzazioni personalizzate in modo uniforme attraverso le nostre applicazioni web, desktop e mobile, ognuna con la propria interfaccia e considerazioni sull'esperienza utente. Questo è stato particolarmente difficile per le nostre app mobili, che dovevano nascondere e mostrare dinamicamente le funzionalità in base alle autorizzazioni dell'utente. Abbiamo affrontato questo problema centralizzando la nostra logica di autorizzazione nel layer API, garantendo che tutte le piattaforme ricevessero dati di autorizzazione coerenti.

Successivamente, abbiamo sviluppato un framework UI flessibile che potesse adattarsi a queste modifiche di autorizzazione in tempo reale, fornendo un'esperienza senza soluzione di continuità indipendentemente dalla piattaforma utilizzata.

L'educazione e l'adozione degli utenti hanno presentato l'ultimo ostacolo nel nostro viaggio delle autorizzazioni personalizzate. Introdurre una funzionalità così potente significava che dovevamo aiutare i nostri utenti a comprendere e sfruttare efficacemente le autorizzazioni personalizzate.

Inizialmente abbiamo lanciato le autorizzazioni personalizzate a un sottoinsieme della nostra base utenti, monitorando attentamente le loro esperienze e raccogliendo informazioni. Questo approccio ci ha permesso di affinare la funzionalità e i nostri materiali educativi in base all'uso reale prima di lanciare a tutta la nostra base utenti.

Il lancio graduale si è rivelato prezioso, aiutandoci a identificare e affrontare problemi minori e punti di confusione degli utenti che non avevamo previsto, portando infine a una funzionalità più rifinita e user-friendly per tutti i nostri utenti.

Questo approccio di lanciare a un sottoinsieme di utenti, così come il nostro tipico periodo di "Beta" di 2-3 settimane sulla nostra Beta pubblica, ci aiuta a dormire sonni tranquilli. :)

## Guardando Avanti

Come per tutte le funzionalità, nulla è mai *"finito"*.

La nostra visione a lungo termine per la funzionalità delle autorizzazioni personalizzate si estende a tag, filtri di campi personalizzati, navigazione di progetto personalizzabile e controlli sui commenti.

Esploriamo ciascun aspetto.

### Autorizzazioni sui Tag

Pensiamo che sarebbe fantastico poter creare autorizzazioni basate sul fatto che un record abbia uno o più tag. Il caso d'uso più ovvio sarebbe quello di creare un ruolo utente personalizzato chiamato "Clienti" e consentire solo agli utenti in quel ruolo di vedere i record che hanno il tag "Clienti".

Questo ti offre una visione immediata di se un record può o meno essere visto dai tuoi clienti.

Questo potrebbe diventare ancora più potente con combinatori AND/OR, dove puoi specificare regole più complesse. Ad esempio, potresti impostare una regola che consente l'accesso a record contrassegnati sia come "Clienti" CHE "Pubblici", o record contrassegnati come "Interni" O "Riservati". Questo livello di flessibilità consentirebbe impostazioni di autorizzazione incredibilmente sfumate, adatte anche alle strutture e ai flussi di lavoro organizzativi più complessi.

Le applicazioni potenziali sono vaste. I project manager potrebbero facilmente segregare informazioni sensibili, i team di vendita potrebbero avere accesso automatico ai dati dei clienti pertinenti e i collaboratori esterni potrebbero essere integrati senza problemi in parti specifiche di un progetto senza rischiare di esporre informazioni interne sensibili.

### Filtri di Campi Personalizzati

La nostra visione per i Filtri di Campi Personalizzati rappresenta un significativo passo avanti nel controllo dell'accesso granulare. Questa funzionalità consentirà agli amministratori di progetto di definire quali record specifici gruppi di utenti possono vedere in base ai valori dei campi personalizzati. Si tratta di creare confini dinamici e basati sui dati per l'accesso alle informazioni.

Immagina di poter impostare autorizzazioni in questo modo:

- Mostra solo i record in cui il menu a discesa "Stato del Progetto" è impostato su "Pubblico"
- Limita la visibilità agli elementi in cui il campo multi-selezione "Dipartimento" include "Marketing"
- Consenti l'accesso ai compiti in cui la casella di controllo "Priorità" è selezionata
- Visualizza i progetti in cui il campo numerico "Budget" è sopra una certa soglia

### Navigazione di Progetto Personalizzabile

Questa è semplicemente un'estensione degli interruttori che abbiamo già. Invece di avere solo interruttori per "attività" e "moduli", vogliamo estendere questo a ogni singola parte della navigazione del progetto. In questo modo, gli amministratori di progetto possono creare interfacce focalizzate e rimuovere strumenti di cui non hanno bisogno.

### Controlli sui Commenti

In futuro, vogliamo essere creativi nel modo in cui consentiamo ai nostri clienti di decidere chi può e non può vedere i commenti. Potremmo consentire aree di commento multiple con schede sotto un record, e ognuna può essere visibile o non visibile a diversi gruppi di utenti.

Inoltre, potremmo anche consentire una funzionalità in cui solo i commenti in cui un utente è *specificamente* menzionato sono visibili, e nient'altro lo è. Questo consentirebbe ai team che hanno clienti sui progetti di garantire che solo i commenti che vogliono che i clienti vedano siano visibili.

## Conclusione

Ecco come abbiamo affrontato la costruzione di una delle funzionalità più interessanti e potenti! [Come puoi vedere nel nostro strumento di confronto per la gestione dei progetti](/compare), pochissimi sistemi di gestione dei progetti hanno una configurazione della matrice delle autorizzazioni così potente, e quelli che lo fanno lo riservano ai loro piani enterprise più costosi, rendendolo inaccessibile a una tipica piccola o media azienda.

Con Blue, hai *tutte* le funzionalità disponibili con il nostro piano: non crediamo che le funzionalità di livello enterprise debbano essere riservate ai clienti enterprise!