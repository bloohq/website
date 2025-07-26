---
title: FAQ sulla Sicurezza di Blue
description: Questa è una lista delle domande più frequenti sui protocolli e le pratiche di sicurezza di Blue.
category: "FAQ"
date: 2024-07-19
---


La nostra missione è organizzare il lavoro nel mondo costruendo la migliore piattaforma di gestione dei progetti del pianeta.

Al centro del raggiungimento di questa missione c'è la garanzia che la nostra piattaforma sia sicura, affidabile e degna di fiducia. Comprendiamo che per essere la tua unica fonte di verità, Blue deve proteggere i tuoi dati aziendali sensibili da minacce esterne, perdita di dati e tempi di inattività.

Questo significa che prendiamo sul serio la sicurezza in Blue.

Quando pensiamo alla sicurezza, consideriamo un approccio olistico che si concentra su tre aree chiave:

1.  **Sicurezza dell'Infrastruttura e della Rete**: Garantisce che i nostri sistemi fisici e virtuali siano protetti da minacce esterne e accessi non autorizzati.
2.  **Sicurezza del Software**: Si concentra sulla sicurezza del codice stesso, comprese le pratiche di codifica sicura, le revisioni regolari del codice e la gestione delle vulnerabilità.
3.  **Sicurezza della Piattaforma**: Include le funzionalità all'interno di Blue, come [controlli di accesso sofisticati](/platform/features/user-permissions), garantendo che i progetti siano privati per impostazione predefinita, e altre misure per proteggere i dati e la privacy degli utenti.

## Quanto è scalabile Blue?

Questa è una domanda importante, poiché desideri un sistema che possa *crescere* con te. Non vuoi dover cambiare la tua piattaforma di gestione dei progetti e dei processi tra sei o dodici mesi.

Scegliamo i fornitori di piattaforme con attenzione, per garantire che possano gestire i carichi di lavoro impegnativi dei nostri clienti. Utilizziamo servizi cloud di alcuni dei migliori fornitori di cloud al mondo che alimentano aziende come [Spotify](https://spotify.com) e [Netflix](https://netflix.com), che hanno un traffico di ordini di grandezza superiore al nostro.

I principali fornitori di cloud che utilizziamo sono:

- **[Cloudflare](https://cloudflare.com)**: Gestiamo il DNS (Domain Name Service) tramite Cloudflare, così come il nostro sito web di marketing che funziona su [Cloudflare Pages](https://pages.cloudflare.com/).
- **[Amazon Web Services](https://aws.amazon.com/)**: Utilizziamo AWS per il nostro database, che è [Aurora](https://aws.amazon.com/rds/aurora/), per l'archiviazione dei file tramite [Simple Storage Service (S3)](https://aws.amazon.com/s3/), e anche per l'invio di email tramite [Simple Email Service (SES)](https://aws.amazon.com/ses/)
- **[Render](https://render.com)**: Utilizziamo Render per i nostri server front-end, server applicazione/API, i nostri servizi in background, sistema di coda e database Redis. Interessante notare che Render è effettivamente costruito *sopra* AWS!

## Quanto sono sicuri i file in Blue?

Iniziamo con l'archiviazione dei dati. I nostri file sono ospitati su [AWS S3](https://aws.amazon.com/s3/), che è il servizio di archiviazione oggetti cloud più popolare al mondo con scalabilità, disponibilità, sicurezza e prestazioni leader del settore.

Abbiamo una disponibilità dei file del 99,99% e una durabilità alta del 99,999999999%.

Analizziamo cosa significa questo.

La disponibilità si riferisce alla quantità di tempo in cui i dati sono operativi e accessibili. La disponibilità dei file del 99,99% significa che possiamo aspettarci che i file siano non disponibili per non più di circa 8,76 ore all'anno.

La durabilità si riferisce alla probabilità che i dati rimangano intatti e non corrotti nel tempo. Questo livello di durabilità significa che possiamo aspettarci di perdere non più di un file su 10 miliardi di file caricati, grazie a una vasta ridondanza e replicazione dei dati tra più data center.

Utilizziamo [S3 Intelligent-Tiering](https://aws.amazon.com/s3/storage-classes/intelligent-tiering/) per spostare automaticamente i file in diverse classi di archiviazione in base alla frequenza di accesso. Basandoci sui modelli di attività di centinaia di migliaia di progetti, notiamo che la maggior parte dei file viene accessibile in un modello che assomiglia a una curva di backoff esponenziale. Questo significa che la maggior parte dei file viene accessibile molto frequentemente nei primi giorni, e poi viene accessibile sempre meno frequentemente. Questo ci consente di spostare i file più vecchi in archiviazioni più lente, ma significativamente più economiche, senza influenzare in modo significativo l'esperienza dell'utente.

I risparmi sui costi per questo sono significativi. S3 Standard-Infrequent Access (S3 Standard-IA) è circa 1,84 volte più economico di S3 Standard. Questo significa che per ogni dollaro che avremmo speso su S3 Standard, spendiamo solo circa 54 centesimi su S3 Standard-IA per la stessa quantità di dati archiviati.

| Caratteristica           | S3 Standard             | S3 Standard-IA       |
|--------------------------|-------------------------|-----------------------|
| Costo di Archiviazione    | $0.023 - $0.021 per GB  | $0.0125 per GB        |
| Costo di Richiesta (PUT, ecc.) | $0.005 per 1.000 richieste | $0.01 per 1.000 richieste |
| Costo di Richiesta (GET) | $0.0004 per 1.000 richieste | $0.001 per 1.000 richieste |
| Costo di Recupero Dati   | $0.00 per GB            | $0.01 per GB          |

I file che carichi tramite Blue sono crittografati sia in transito che a riposo. I dati trasferiti da e verso Amazon S3 sono protetti utilizzando [Transport Layer Security (TLS)](https://www.internetsociety.org/deploy360/tls/basics), proteggendo contro [intercettazioni](https://en.wikipedia.org/wiki/Network_eavesdropping) e [attacchi man-in-the-middle](https://en.wikipedia.org/wiki/Man-in-the-middle_attack). Per la crittografia a riposo, Amazon S3 utilizza la crittografia lato server (SSE-S3), che crittografa automaticamente tutti i nuovi caricamenti con crittografia AES-256, con Amazon che gestisce le chiavi di crittografia. Questo garantisce che i tuoi dati rimangano sicuri durante l'intero ciclo di vita.

## E per i dati non file?

Il nostro database è alimentato da [AWS Aurora](https://aws.amazon.com/rds/aurora/), un moderno servizio di database relazionale che garantisce alte prestazioni, disponibilità e sicurezza per i tuoi dati.

I dati in Aurora sono crittografati sia in transito che a riposo. Utilizziamo SSL (AES-256) per proteggere le connessioni tra la tua istanza di database e la tua applicazione, proteggendo i dati durante il trasferimento. Per la crittografia a riposo, Aurora utilizza chiavi gestite tramite AWS Key Management Service (KMS), garantendo che tutti i dati memorizzati, comprese le copie di sicurezza automatiche, gli snapshot e le repliche, siano crittografati e protetti.

Aurora presenta un sistema di archiviazione distribuito, tollerante ai guasti e auto-riparabile. Questo sistema è decoupled dalle risorse di calcolo e può scalare automaticamente fino a 128 TiB per istanza di database. I dati sono replicati in tre [Availability Zones](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) (AZ), fornendo resilienza contro la perdita di dati e garantendo alta disponibilità. In caso di un crash del database, Aurora riduce i tempi di recupero a meno di 60 secondi, garantendo una minima interruzione.

Blue esegue continuamente il backup del nostro database su Amazon S3, abilitando il recupero a un punto nel tempo. Questo significa che possiamo ripristinare il database master di Blue a qualsiasi momento specifico negli ultimi cinque minuti, garantendo che i tuoi dati siano sempre recuperabili. Effettuiamo anche snapshot regolari del database per periodi di retention più lunghi.

Essendo un servizio completamente gestito, Aurora automatizza compiti di amministrazione che richiedono tempo, come la fornitura di hardware, la configurazione del database, le patch e i backup. Questo riduce il carico operativo e garantisce che il nostro database sia sempre aggiornato con le ultime patch di sicurezza e miglioramenti delle prestazioni.

Se siamo più efficienti, possiamo trasferire i nostri risparmi sui costi ai nostri clienti con la nostra [tariffazione leader del settore](/pricing).

Aurora è conforme a vari standard del settore come HIPAA, GDPR e SOC 2, garantendo che le tue pratiche di gestione dei dati soddisfino requisiti normativi rigorosi. Audit di sicurezza regolari e integrazione con [Amazon GuardDuty](https://aws.amazon.com/guardduty/) aiutano a rilevare e mitigare potenziali minacce alla sicurezza.

## Come garantisce Blue la sicurezza del login?

Blue utilizza [link magici via email](https://documentation.blue.cc/user-management/magic-links) per fornire accesso sicuro e conveniente al tuo account, eliminando la necessità di password tradizionali.

Questo approccio migliora significativamente la sicurezza mitigando le minacce comuni associate ai login basati su password. Eliminando le password, i link magici proteggono contro attacchi di phishing e furto di password, *poiché non c'è password da rubare o sfruttare.*

Ogni link magico è valido per una sola sessione di login, riducendo il rischio di accesso non autorizzato. Inoltre, questi link scadono dopo 15 minuti, garantendo che eventuali link non utilizzati non possano essere sfruttati, migliorando ulteriormente la sicurezza.

La comodità offerta dai link magici è anche notevole. I link magici forniscono un'esperienza di login senza problemi, consentendoti di accedere al tuo account *senza* dover ricordare password complesse.

Questo non solo semplifica il processo di login, ma previene anche le violazioni della sicurezza che si verificano quando le password vengono riutilizzate su più servizi. Molti utenti tendono a utilizzare la stessa password su varie piattaforme, il che significa che una violazione della sicurezza su un servizio potrebbe compromettere i loro account su altri servizi, incluso Blue. Utilizzando i link magici, la sicurezza di Blue non dipende dalle pratiche di sicurezza di altri servizi, fornendo un livello di protezione più robusto e indipendente per i nostri utenti.

Quando richiedi di accedere al tuo account Blue, un URL di login unico viene inviato alla tua email. Cliccando su questo link ti registrerai istantaneamente nel tuo account. Il link è progettato per scadere dopo un singolo utilizzo o dopo 15 minuti, a seconda di quale evento si verifica per primo, aggiungendo un ulteriore livello di sicurezza. Utilizzando i link magici, Blue garantisce che il tuo processo di login sia sia sicuro che user-friendly, fornendo tranquillità e convenienza.

## Come posso controllare l'affidabilità e il tempo di attività di Blue?

In Blue, ci impegniamo a mantenere un alto livello di affidabilità e trasparenza per i nostri utenti. Per fornire visibilità sulle prestazioni della nostra piattaforma, offriamo una [pagina di stato del sistema dedicata](https://status.blue.cc) che è anche collegata dal nostro footer su ogni pagina del nostro sito web.

![](/insights/status-page.png)

Questa pagina mostra i nostri dati storici di uptime, consentendoti di vedere quanto costantemente i nostri servizi siano stati disponibili nel tempo. Inoltre, la pagina di stato include rapporti dettagliati sugli incidenti, fornendo trasparenza su eventuali problemi passati, il loro impatto e i passi che abbiamo intrapreso per risolverli e prevenire future occorrenze.