/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

import { defineConfig } from 'vite'
import solid from 'vite-plugin-solid'
import suidPlugin from "@suid/vite-plugin"

export default defineConfig({
	plugins: [ solid(), suidPlugin() ]
})
