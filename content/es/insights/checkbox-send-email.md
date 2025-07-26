---
title: Automatización de gestión de proyectos — correos electrónicos a las partes interesadas.
category: "Product Updates"
description: A menudo, usted quiere tener el control de sus automatizaciones de gestión de proyectos
date: 2024-07-08
---

Hemos cubierto cómo [crear automatizaciones de correo electrónico anteriormente.](/insights/email-automations)

Sin embargo, a menudo hay partes interesadas en los proyectos que solo necesitan ser alertadas cuando hay algo *realmente* importante.

¿No sería genial si hubiera una automatización de gestión de proyectos donde usted, como gerente de proyecto, pudiera tener el control de *exactamente* cuándo notificar a una parte interesada clave con solo presionar un botón?

Bueno, resulta que con Blue, ¡puede hacer precisamente esto!

Hoy vamos a aprender cómo crear una automatización de gestión de proyectos realmente útil:

Una casilla de verificación que notifica automáticamente a una o más partes interesadas clave, brindándoles todo el contexto clave de lo que les está notificando. Como punto adicional, también aprenderemos cómo restringir esta capacidad para que solo ciertos miembros de su proyecto puedan activar esta notificación por correo electrónico.

Esto se verá algo así una vez que haya terminado:

![](/insights/checkbox-email-automation.png)

Y solo marcando esta casilla de verificación, podrá activar una automatización de gestión de proyectos para enviar un correo electrónico de notificación personalizado a las partes interesadas.

Vamos paso a paso.

## 1. Cree su campo personalizado de casilla de verificación

Esto es muy fácil, puede consultar nuestra [documentación detallada](https://documentation.blue.cc/custom-fields/introduction#creating-custom-fields) sobre la creación de campos personalizados.

Asegúrese de nombrar este campo con algo obvio que recordará, como "notificar a la gerencia" o "notificar a las partes interesadas".

## 2. Cree su disparador de automatización de gestión de proyectos.

En la vista de registros de su proyecto, haga clic en el pequeño robot en la parte superior derecha para abrir la configuración de automatización:

<video autoplay loop muted playsinline>
  <source src="/videos/notify-stakeholders-automation-setup.mp4" type="video/mp4">
</video>

## 3. Cree su acción de automatización de gestión de proyectos.

En este caso, nuestra acción será enviar una notificación personalizada por correo electrónico a una o más direcciones de correo electrónico. Es bueno tener en cuenta aquí que estas personas **no** tienen que estar en Blue para recibir estos correos electrónicos, puede enviar correos electrónicos a *cualquier* dirección de correo electrónico.

Puede obtener más información en nuestra [guía de documentación detallada sobre cómo configurar automatizaciones de correo electrónico](https://documentation.blue.cc/automations/actions/email-automations)

Su resultado final debería verse algo así:

![](/insights/email-automation-example.png)

## 4. Bonus: Restrinja el acceso a la casilla de verificación.

Puede usar [roles de usuario personalizados en Blue](/platform/features/user-permissions) para restringir el acceso a los campos personalizados de casilla de verificación, asegurando que solo los miembros autorizados del equipo puedan activar las notificaciones por correo electrónico.

Blue permite a los Administradores de Proyecto definir roles y asignar permisos a grupos de usuarios. Este sistema es crucial para mantener el control sobre quién puede interactuar con elementos específicos de su proyecto, incluidos los campos personalizados como la casilla de verificación de notificación.

1. Navegue a la sección de Gestión de Usuarios en Blue y seleccione "Roles de Usuario Personalizados".
2. Cree un nuevo rol proporcionando un nombre descriptivo y una descripción opcional.
3. Dentro de los permisos del rol, localice la sección para Acceso a Campos Personalizados.
4. Especifique si el rol puede ver o editar el campo personalizado de casilla de verificación. Por ejemplo, restrinja el acceso de edición a roles como "Administrador de Proyecto" mientras permite que un rol personalizado recién creado administre este campo.
5. Asigne el rol recién creado a los usuarios o grupos de usuarios apropiados. Esto garantiza que solo las personas designadas tengan la capacidad de interactuar con la casilla de verificación de notificación.

[Lea más en nuestro sitio de documentación oficial.](https://documentation.blue.cc/user-management/roles/custom-user-roles)

Al implementar estos roles personalizados, usted mejora la seguridad e integridad de sus procesos de gestión de proyectos. Solo los miembros autorizados del equipo pueden activar notificaciones críticas por correo electrónico, asegurando que las partes interesadas reciban actualizaciones importantes sin alertas innecesarias.

## Conclusión

Al implementar la automatización de gestión de proyectos descrita anteriormente, usted obtiene un control preciso sobre cuándo y cómo notificar a las partes interesadas clave. Este enfoque garantiza que las actualizaciones importantes se comuniquen de manera efectiva, sin abrumar a sus partes interesadas con información innecesaria. Utilizando las funciones de campos personalizados y automatización de Blue, puede optimizar su proceso de gestión de proyectos, mejorar la comunicación y mantener un alto nivel de eficiencia.

Con solo una simple casilla de verificación, puede activar notificaciones personalizadas por correo electrónico adaptadas a las necesidades de su proyecto, asegurando que las personas correctas estén informadas en el momento correcto. Además, la capacidad de restringir esta funcionalidad a miembros específicos del equipo agrega una capa adicional de control y seguridad.

Comience a aprovechar esta poderosa función en Blue hoy para mantener informadas a sus partes interesadas y sus proyectos funcionando sin problemas. Para obtener pasos más detallados y opciones de personalización adicionales, consulte los enlaces de documentación proporcionados. ¡Feliz automatización!