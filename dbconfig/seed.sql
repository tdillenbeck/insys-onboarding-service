-- How to run:
-- psql -U postgres -h localhost -p 5432 -d insys_onboarding_local -f dbconfig/seed.sql

-- NOTE: The Task ID's are used in the creating task instances. Make sure to update the ID's in psql/taskinstance.go

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
-- Software Tasks
  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '16a6dc91-ec6b-4b09-b591-a5b0dfa92932', -- id
    'Sync your patient data to Weave', -- title
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> On this scheduled call, we will help get your patient or customer database syncing with Weave. This will allow you to start using many of Weave''s software features.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30-60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> We may need to access your office server - so make sure you have access to that computer and the right login information.</div>',
    2, -- display_order
    default, -- created_at
    default, -- updated_at
    '26ba2237-c452-42dd-95ca-a5e59dd2853b', -- onboarding_category_id
    'Schedule Call', -- button_content
    'https://getweave.chilipiper.com/book/installation-scheduler?type=software-installation' -- button_external_url
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Sync your patient data to Weave', -- title
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> On this scheduled call, we will help get your patient or customer database syncing with Weave. This will allow you to start using many of Weave''s software features.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30-60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> We may need to access your office server - so make sure you have access to that computer and the right login information.</div>',
    2, -- display_order
    default, -- created_at
    default, -- updated_at
    'Schedule Call', -- button_content
    'https://getweave.chilipiper.com/book/installation-scheduler?type=software-installation' -- button_external_url
  );


  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '720af494-38a4-499f-8633-9c8d5169cd43', -- id
    'Install Weave on other workstations in your office', -- title
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of Weave on other workstations in your office. Installing Weave on more workstations will help you make the most of features like Team Chat and two-way texting.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 5 minutes per workstation</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Use the link below to download Weave throughout your office - if you need help, just let us know!</div>',
    5, -- display_order
    default, -- created_at
    default, -- updated_at
    '26ba2237-c452-42dd-95ca-a5e59dd2853b', -- onboarding_category_id
    'Install Weave', --button_content
    'http://www.weavehelp.com/software-install' -- button_external_url
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Install Weave on other workstations in your office', -- title
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of Weave on other workstations in your office. Installing Weave on more workstations will help you make the most of features like Team Chat and two-way texting.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 5 minutes per workstation</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Use the link below to download Weave throughout your office - if you need help, just let us know!</div>',
    5, -- display_order
    default, -- created_at
    default, -- updated_at
    'Install Weave', --button_content
    'http://www.weavehelp.com/software-install' -- button_external_url
  );

  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    'c20b65d8-e281-4e62-98f0-4aebf83e0bee',
    'Watch our helpful software training videos',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features you can be using now, like two-way texting, Team Chat, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>',
    6,
    default,
    default,
    '26ba2237-c452-42dd-95ca-a5e59dd2853b',
    'Watch Videos',
    'http://www.weavehelp.com/webinar-on-demand/'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Watch our helpful software training videos',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features you can be using now, like two-way texting, Team Chat, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>',
    6,
    default,
    default,
    'Watch Videos',
    'http://www.weavehelp.com/webinar-on-demand/'
  );


  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '1120842a-a24b-40e8-b29e-0e05e89af99f',
    'Install the Weave mobile app',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of our mobile app on your phone. Our mobile app is a great tool to see your schedule, field office calls, and chat with your team - all on-the-go, whenever you need access.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 3-5 minutes</div>',
    7,
    default,
    default,
    '26ba2237-c452-42dd-95ca-a5e59dd2853b',
    'Learn How',
    'http://www.weavehelp.com/weave-mobile-app/'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Install the Weave mobile app',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A quick installation of our mobile app on your phone. Our mobile app is a great tool to see your schedule, field office calls, and chat with your team - all on-the-go, whenever you need access.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 3-5 minutes</div>',
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
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will remotely access your workstation to check your office network and make recommendations to have the best experience with Weave.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 15 minutes</div>',
    3,
    default,
    default,
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Schedule Call',
    'https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=5221365'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Check your office network to ensure compatibility',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will remotely access your workstation to check your office network and make recommendations to have the best experience with Weave.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 15 minutes</div>',
    3,
    default,
    default,
    'Schedule Call',
    'https://app.acuityscheduling.com/schedule.php?owner=14911380&appointmentType=5221365'
  );


  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    'fd4f656c-c9f1-47b8-96ad-3080b999a843',
    'Install your new phones',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will work with you (or your IT professional) to guide you through physically connecting the new phones to internet and power.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Typically, we will want to help you connect the new phones side-by-side with the old phones while we work with your current phone company for a few days to transition your phones service to Weave.</div>',
    8,
    default,
    default,
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Schedule Call',
    'https://getweave.chilipiper.com/book/installation-scheduler?type=phone-installation'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Install your new phones',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which a Weave technician will work with you (or your IT professional) to guide you through physically connecting the new phones to internet and power.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 60 minutes</div><div class="insys-content-body"><span class="insys-content-bold">Anything else?</span> Typically, we will want to help you connect the new phones side-by-side with the old phones while we work with your current phone company for a few days to transition your phones service to Weave.</div>',
    8,
    default,
    default,
    'Schedule Call',
    'https://getweave.chilipiper.com/book/installation-scheduler?type=phone-installation'
  );


  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '47743fae-c775-45d5-8a51-dc7e3371dfa4',
    'Watch our helpful phone training videos',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features of your new phone system, including voicemail and auto-attendant setup, transferring calls, placing callers on hold, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>',
    9,
    default,
    default,
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Watch Videos',
    'http://www.weavehelp.com/webinar-on-demand/'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Watch our helpful phone training videos',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> An online webinar - available to schedule or on demand - that will guide you through all the key features of your new phone system, including voicemail and auto-attendant setup, transferring calls, placing callers on hold, and much more.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>',
    9,
    default,
    default,
    'Watch Videos',
    'http://www.weavehelp.com/webinar-on-demand/'
  );


  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '2d2df285-9211-48fc-a057-74f7dee2d9a4',
    'Customize your phone system',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which your onboarding agent will help customize various features of your phone system - like which phones ring on inbound calls, the name and extension of each phone, and any advanced call routing or auto attendant.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>',
    10,
    default,
    default,
    'ebc72a11-f1b3-40d5-888e-5b6aba66e871',
    'Schedule Call',
    'https://getweave.chilipiper.com/book/customization-calls'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url) = (
    'Customize your phone system',
    '<div class="insys-content-body"><span class="insys-content-bold">What is this?</span> A scheduled call during which your onboarding agent will help customize various features of your phone system - like which phones ring on inbound calls, the name and extension of each phone, and any advanced call routing or auto attendant.</div><div class="insys-content-body"><span class="insys-content-bold">How long will this take?</span> 30 minutes</div>',
    10,
    default,
    default,
    'Schedule Call',
    'https://getweave.chilipiper.com/book/customization-calls'
  );


-- Porting Tasks
  INSERT INTO insys_onboarding.onboarding_tasks VALUES (
    '9aec502b-f8b8-4f10-9748-1fe4050eacde',
    'Provide current phone account info',
    '<div class="insys-content-body"> <span class="insys-content-bold"> What is this?  </span> In order to port your phone numbers from your current provider to your new Weave phones, we need three things: <ul> <li>Info about your current phone account</li> <li>Your approval of the terms of service</li> <li>A copy of your current phone bill</li> </ul> </div> <div class="insys-content-body"> <span class="insys-content-bold">How long will this take?</span> 5-10 minutes </div>',
    4,
    default,
    default,
    'd0da53a9-fbdb-4d22-85c6-ed521f237349',
    'Let''s Start',
    'https://getweave.chilipiper.com/book/start-porting-process-call',
    '/porting'
  )
  ON CONFLICT(id) DO UPDATE SET (title, content, display_order, created_at, updated_at, button_content, button_external_url, button_internal_url) = (
    'Provide current phone account info',
    '<div class="insys-content-body"> <span class="insys-content-bold"> What is this?  </span> In order to port your phone numbers from your current provider to your new Weave phones, we need three things: <ul> <li>Info about your current phone account</li> <li>Your approval of the terms of service</li> <li>A copy of your current phone bill</li> </ul> </div> <div class="insys-content-body"> <span class="insys-content-bold">How long will this take?</span> 5-10 minutes </div>',
    4,
    default,
    default,
    'Let''s Start',
    'https://getweave.chilipiper.com/book/start-porting-process-call',
    '/porting'
  );
