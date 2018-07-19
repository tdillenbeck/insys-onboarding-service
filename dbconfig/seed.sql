-- How to run:
-- psql -U postgres -h localhost -p 5432 -d insys_onboarding_local -f dbconfig/seed.sql

-- Categories
-- INSERT INTO insys_onboarding.onboarding_categories VALUES (id, display_text, display_order, created_at, updated_at);
  INSERT INTO insys_onboarding.onboarding_categories VALUES (
    '26ba2237-c452-42dd-95ca-a5e59dd2853b',
    'Software',
    1,
    default,
    default
  )
  ON CONFLICT(id) DO UPDATE SET (display_text, display_order, updated_at) = (
    'Software',
    1,
    default
  );

  INSERT INTO insys_onboarding.onboarding_categories VALUES (
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Phones',
    2,
    default,
    default
  )
  ON CONFLICT(id) DO UPDATE SET (display_text, display_order, updated_at) = (
    'Phones',
    2,
    default
  );

  INSERT INTO insys_onboarding.onboarding_categories VALUES (
    'd0da53a9-fbdb-4d22-85c6-ed521f237349',
    'Porting',
    3,
    default,
    default
  )
  ON CONFLICT(id) DO UPDATE SET (display_text, display_order, updated_at) = (
    'Porting',
    3,
    default
  );

-- Tasks
-- INSERT INTO insys_onboarding.onboarding_tasks VALUES (id, title, content, display_order, created_at, updated_at, onboarding_category_id, button_content, button_external_url);

-- Software Tasks
  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '16a6dc91-ec6b-4b09-b591-a5b0dfa92932', -- id
    'Sync your patient data to Weave', -- title
    'In order to start using all our software features, we need to sync your patient data to Weave. Typically, we will do this by installing a sync application on your office server - so make sure you have access to that computer and the right login information. If your practice management system is cloud-based, we will instead walk you through how to generate the credentials we need. Click below to schedule a 30-60 minute phone call with a Weave technician.', -- content
    2, -- display_order
    default, -- created_at
    default, -- updated_at
    '26ba2237-c452-42dd-95ca-a5e59dd2853b', -- onboarding_category_id
    'Schedule Call', -- button_content
    'https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=5221365' -- button_external_url
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Sync your patient data to Weave', -- title
    'In order to start using all our software features, we need to sync your patient data to Weave. Typically, we will do this by installing a sync application on your office server - so make sure you have access to that computer and the right login information. If your practice management system is cloud-based, we will instead walk you through how to generate the credentials we need. Click below to schedule a 30-60 minute phone call with a Weave technician.', -- content
    2, -- display_order
    default, -- created_at
    default, -- updated_at
    'Schedule Call', -- button_content
    'https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=5221365' -- button_external_url
  );

  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '720af494-38a4-499f-8633-9c8d5169cd43', -- id
    'Install Weave on other workstations in your office', -- title
    'In order to make the most of features like Team Chat and two-way text messaging, we recommend installing Weave on most workstations throughout your office. Click below to find helpful instructions to download Weave throughout your office.', -- content
    5, -- display_order
    default, -- created_at
    default, -- updated_at
    '26ba2237-c452-42dd-95ca-a5e59dd2853b', -- onboarding_category_id
    'Install Weave', --button_content
    'http://www.weavehelp.com/software-install' -- button_external_url
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Install Weave on other workstations in your office', -- title
    'In order to make the most of features like Team Chat and two-way text messaging, we recommend installing Weave on most workstations throughout your office. Click below to find helpful instructions to download Weave throughout your office.', -- content
    5, -- display_order
    default, -- created_at
    default, -- updated_at
    'Install Weave', --button_content
    'http://www.weavehelp.com/software-install' -- button_external_url
  );

  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    'c20b65d8-e281-4e62-98f0-4aebf83e0bee',
    'Watch our helpful software training videos',
    'In order to get the most out of the Weave software, we recommend tuning in to our training videos online. Our helpful training staff will walk you through setting up and using two-way texting, automated appointment reminders, and all our other features. Click below to watch our online webinars',
    6,
    default,
    default,
    '26ba2237-c452-42dd-95ca-a5e59dd2853b',
    'Watch Videos',
    'http://www.weavehelp.com/webinar-registration-page/'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Watch our helpful software training videos',
    'In order to get the most out of the Weave software, we recommend tuning in to our training videos online. Our helpful training staff will walk you through setting up and using two-way texting, automated appointment reminders, and all our other features. Click below to watch our online webinars',
    6,
    default,
    default,
    'Watch Videos',
    'http://www.weavehelp.com/webinar-registration-page/'
  );

  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '1120842a-a24b-40e8-b29e-0e05e89af99f',
    'Install the Weave mobile app',
    'In order to see your schedule, field calls, and chat with your team all on the go, we recommend installing the Weave app on your mobile device. Click below to find helpful instructions to download Weave on your phone.',
    7,
    default,
    default,
    '26ba2237-c452-42dd-95ca-a5e59dd2853b',
    'Learn How',
    'http://www.weavehelp.com/weave-mobile-app/'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Install the Weave mobile app',
    'In order to see your schedule, field calls, and chat with your team all on the go, we recommend installing the Weave app on your mobile device. Click below to find helpful instructions to download Weave on your phone.',
    7,
    default,
    default,
    'Learn How',
    'http://www.weavehelp.com/weave-mobile-app/'
  );

-- Phone Tasks
  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '7b15e061-8002-4edc-9bf4-f38c6eec6364',
    'Check your office network to ensure compatibility',
    'In order to ensure the highest quality phone service on Weave''s internet-based phones, we need to verify your network compatibility. Click below to schedule a 15 minute phone call with a Weave technician. <a href="http://www.weavehelp.com/network-specs">Learn More</a>',
    3,
    default,
    default,
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Schedule Call',
    'https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=5221365'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Check your office network to ensure compatibility',
    'In order to ensure the highest quality phone service on Weave''s internet-based phones, we need to verify your network compatibility. Click below to schedule a 15 minute phone call with a Weave technician. <a href="http://www.weavehelp.com/network-specs">Learn More</a>',
    3,
    default,
    default,
    'Schedule Call',
    'https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=5221365'
  );

  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    'fd4f656c-c9f1-47b8-96ad-3080b999a843',
    'Install your new phones',
    'You''re making great progress! Click below to schedule a 60 minute phone call with a Weave technician who will guide you through the process of connecting the new phones. Typically, we will aim to help you plug the new phones in side-by-side with your old phones. After finishing, you will continue to use your old phones while we work with your current phone service provider for a few days to officially move your phone numbers to our system. <a href="http://www.weavehelp.com/phone-install">Learn more.</a>',
    8,
    default,
    default,
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Schedule Call',
    'https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=category:Installs'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Install your new phones',
    'You''re making great progress! Click below to schedule a 60 minute phone call with a Weave technician who will guide you through the process of connecting the new phones. Typically, we will aim to help you plug the new phones in side-by-side with your old phones. After finishing, you will continue to use your old phones while we work with your current phone service provider for a few days to officially move your phone numbers to our system. <a href="http://www.weavehelp.com/phone-install">Learn more.</a>',
    8,
    default,
    default,
    'Schedule Call',
    'https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=category:Installs'
  );

  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '47743fae-c775-45d5-8a51-dc7e3371dfa4',
    'Watch our helpful phone training videos',
    'In order to get the most out of the Weave phones, we recommend tuning in to our training videos online. Our helpful training staff will walk you through setting up voicemail messages, placing callers on hold, transferring calls, and all our other phone features. Click below to watch our online webinars.',
    9,
    default,
    default,
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Watch Videos',
    'http://www.weavehelp.com/webinar-registration-page/'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Watch our helpful phone training videos',
    'In order to get the most out of the Weave phones, we recommend tuning in to our training videos online. Our helpful training staff will walk you through setting up voicemail messages, placing callers on hold, transferring calls, and all our other phone features. Click below to watch our online webinars.',
    9,
    default,
    default,
    'Watch Videos',
    'http://www.weavehelp.com/webinar-registration-page/'
  );

  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '2d2df285-9211-48fc-a057-74f7dee2d9a4',
    'Customize your phone system',
    'In order to have the best experience with your new phones, we recommend customizing the system to suit your needs. Your onboarding agent will help you change which phones ring when a call comes in, adjust the names and extensions on each phone, and set up any advanced call routing or auto attendant you may need. Click below to schedule a call.',
    10,
    default,
    default,
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Schedule Call',
    'https://getweave.chilipiper.com/book/customization-calls'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Customize your phone system',
    'In order to have the best experience with your new phones, we recommend customizing the system to suit your needs. Your onboarding agent will help you change which phones ring when a call comes in, adjust the names and extensions on each phone, and set up any advanced call routing or auto attendant you may need. Click below to schedule a call.',
    10,
    default,
    default,
    'Schedule Call',
    'https://getweave.chilipiper.com/book/customization-calls'
  );

-- Porting Tasks
  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '9aec502b-f8b8-4f10-9748-1fe4050eacde',
    'Verify your phone numbers',
    'One of your first steps is to work with your onboarding agent to verify your office phone numbers. Your onboarding agent will then handle some initial processing with your current phone service provider to ensure that when your office is ready to start using the new phones, we are ready to activate them for you. Click below to schedule a call with your onboarding agent.',
    4,
    default,
    default,
    'd0da53a9-fbdb-4d22-85c6-ed521f237349',
    'Schedule Call',
    'https://getweave.chilipiper.com/book/start-porting-process-call'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Verify your phone numbers',
    'One of your first steps is to work with your onboarding agent to verify your office phone numbers. Your onboarding agent will then handle some initial processing with your current phone service provider to ensure that when your office is ready to start using the new phones, we are ready to activate them for you. Click below to schedule a call with your onboarding agent.',
    4,
    default,
    default,
    'Schedule Call',
    'https://getweave.chilipiper.com/book/start-porting-process-call'
  );

