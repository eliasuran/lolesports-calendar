import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
  integrations: [
    starlight({
      title: 'lolesports-calendar',
      social: {
        github: 'https://github.com/eliasuran/lolesports-calendar',
      },
      sidebar: [
        {
          label: 'Getting Started',
          items: [{ label: 'Introduction', link: '/introduction' }],
        },
        {
          label: 'Api',
          autogenerate: { directory: 'api' },
          badge: 'New',
        },
      ],
    }),
  ],
});
