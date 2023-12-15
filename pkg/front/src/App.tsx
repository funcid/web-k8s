/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

import { createSignal } from "solid-js";
import { Alert, Box, Button, Chip, InputLabel, MenuItem, Select, Stack, TextField } from "@suid/material";
import { SelectChangeEvent } from "@suid/material/Select";

enum RestartPolicy { Always = 'Always', OnFailure = 'OnFailure', Never = 'Never' }

enum DeploymentStrategy { Recreate = 'Recreate', RollingUpdate = 'RollingUpdate' }

interface PodRequest {
	name: string,
	namespace: string,
	labels: { [key: string]: string },
	restartPolicy: RestartPolicy,
	containers: {
		name: string
		image: string
		command: string[] | undefined
		args: string[] | undefined
		ports: { containerPort: number }[]
	}[]
}

interface DeploymentRequest {
	name: string,
	namespace: string,
	labels: { [key: string]: string },
	replicas: number,
	strategy: DeploymentStrategy,
	matchLabels: { [key: string]: string },
	pod: PodRequest
}

interface DeleteRequest {
	name: string,
	namespace: string
}

interface Result {
	success: boolean,
	error: string
}

async function makeRequest(
	method: 'POST' | 'DELETE',
	path: string,
	request: unknown,
	setResult: (value: Result) => void
) {
	const res = await fetch(path, {
		method,
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(request),
	})

	setResult(await res.json() as Result)
}

function renderResult(result: Result | undefined) {
	if (result != undefined) {
		return <Alert severity={result.success ? 'success' : 'error'}>{result.error}</Alert>
	}
	return <></>
}

function chipIfNotZero(val: number, label: string) {
	if (val != 0) {
		return <Chip size="small" color="success" label={label} />
	}
}

enum Type {
	Deployment = "Deployment",
	DeleteDeployment = "DeleteDeployment"
}

function undefinedIfBlank(str: string): string | undefined {
	return str.length == 0 ? undefined : str
}

interface ContainerDesc {
	name: string | undefined,
	image: string | undefined,
	command: string[] | undefined,
	args: string[] | undefined,
	ports: { containerPort: number }[]
}

interface LabelDesc {
	key: string | undefined,
	value: string | undefined
}

function labelsToMap(labels: LabelDesc[]): { [key: string]: string } {
	const _labels: { [key: string]: string } = {}

	for (let label of labels) {
		_labels[label.key!] = label.value!
	}

	return _labels
}

function Label(props: { obj: LabelDesc }) {
	return <Box>
		<TextField
			required
			label="Key"
			defaultValue={props.obj.key}
			onChange={(e) => props.obj.key = undefinedIfBlank(e.target.value)}
		/>
		<TextField
			required
			label="Value"
			defaultValue={props.obj.value}
			onChange={(e) => props.obj.value = undefinedIfBlank(e.target.value)}
		/>
	</Box>
}

function Port(props: { container: ContainerDesc, idx: number }) {
	const { container, idx } = props

	return <TextField
		required
		label="Port"
		defaultValue={container.ports[idx]?.containerPort}
		onChange={(e) => {
			let val = undefinedIfBlank(e.target.value)
			if (val == undefined) {
				container.ports[idx] = { containerPort: -props.idx - 1 }
				return
			}
			container.ports[idx] = { containerPort: +val! }
		}}
	/>
}

function Container(props: { obj: ContainerDesc }) {
	let desc = props.obj

	const [ ports$, setPorts ] = createSignal<number>(desc.ports.length)

	return <Box>
		<Stack sx={{ width: '100%', maxWidth: "500px" }} spacing={2}>
			<Chip size="small" label="Container" />

			<TextField
				required
				label="Name"
				defaultValue={desc.name}
				onChange={(e) => desc.name = undefinedIfBlank(e.target.value)}
			/>

			<TextField
				required
				label="Image"
				defaultValue={desc.image}
				onChange={(e) => desc.image = undefinedIfBlank(e.target.value)}
			/>

			<TextField
				label="Command"
				defaultValue={desc.command == undefined ? undefined : desc.command.join(' ')}
				onChange={(e) => {
					const val = undefinedIfBlank(e.target.value)
					if (val == undefined) {
						desc.command = undefined
						return
					}
					desc.command = val.split(" ")
				}}
			/>

			<TextField
				label="Args"
				defaultValue={desc.args == undefined ? undefined : desc.args.join(' ')}
				onChange={(e) => {
					const val = undefinedIfBlank(e.target.value)
					if (val == undefined) {
						desc.args = undefined
						return
					}
					desc.args = val.split(" ")
				}}
			/>

			{chipIfNotZero(ports$(), 'Ports')}
			{[ ...Array(ports$()).keys() ].map(n => <Port container={desc} idx={n} />)}

			<Button onClick={() => {
				setPorts(ports$() + 1)
			}} variant="contained">Add Port</Button>
		</Stack>
	</Box>
}

function Deployment() {
	const [ result$, setResult ] = createSignal<Result>()
	const [ containers$, setContainers ] = createSignal<ContainerDesc[]>([])
	const [ deploymentLabels$, setDeploymentLabels ] = createSignal<LabelDesc[]>([])
	const [ podLabels$, setPodLabels ] = createSignal<LabelDesc[]>([])
	const [ matchLabels$, setMatchLabels ] = createSignal<LabelDesc[]>([])

	let deploymentName: string | undefined = undefined
	let replicas: number = 0
	let strategy: DeploymentStrategy | undefined = undefined
	let namespace: string | undefined = "default"
	let podName: string | undefined = undefined

	function create() {
		let deploymentLabels = deploymentLabels$()
		let podLabels = podLabels$()
		let matchLabels = matchLabels$()
		let containers = containers$()

		if (deploymentName == undefined) {
			setResult({ success: false, error: "deployment name is empty" })
			return
		}

		if (replicas == 0) {
			setResult({ success: false, error: "replicas is 0 / empty" })
			return
		}

		if (strategy == undefined) {
			setResult({ success: false, error: "strategy is empty" })
			return
		}

		if (podName == undefined) {
			setResult({ success: false, error: "pod name is empty" })
			return
		}
		if (namespace == undefined) {
			setResult({ success: false, error: "namespace is empty" })
			return;
		}
		for (let label of podLabels) {
			if (label.key == undefined) {
				setResult({ success: false, error: "pod label label key is empty" });
				return
			}
			if (label.value == undefined) {
				setResult({ success: false, error: "pod label value is empty" });
				return
			}
		}
		for (let label of deploymentLabels) {
			if (label.key == undefined) {
				setResult({ success: false, error: "deployment label label key is empty" });
				return
			}
			if (label.value == undefined) {
				setResult({ success: false, error: "deployment label value is empty" });
				return
			}
		}

		for (let label of matchLabels) {
			if (label.key == undefined) {
				setResult({ success: false, error: "match label label key is empty" });
				return
			}
			if (label.value == undefined) {
				setResult({ success: false, error: "match label value is empty" });
				return
			}
		}

		for (let container of containers) {
			if (container.name == undefined) {
				setResult({ success: false, error: "container name is empty" });
				return
			}
			if (container.image == undefined) {
				setResult({ success: false, error: "container image is empty" });
				return
			}

			for (let { containerPort } of container.ports) {
				if (containerPort < 0) {
					setResult({ success: false, error: `port #${-containerPort} is empty` });
					return;
				}
			}
		}

		const request: DeploymentRequest = {
			name: deploymentName,
			namespace: namespace,
			replicas: replicas,
			strategy: strategy,
			labels: labelsToMap(deploymentLabels),
			matchLabels: labelsToMap(matchLabels),
			pod: {
				name: podName,
				namespace: namespace,
				labels: labelsToMap(podLabels),
				restartPolicy: RestartPolicy.Always,
				containers: containers.map(container => {
					return {
						name: container.name!,
						image: container.image!,
						command: container.command,
						args: container.args,
						ports: container.ports
					}
				})
			}
		};

		void makeRequest('POST', `/v1/deployment`, request, setResult)
	}

	function addContainer() {
		setContainers([ ...containers$(), {
			name: undefined,
			image: undefined,
			command: undefined,
			args: undefined,
			ports: []
		} ])
	}

	function addPodLabel() {
		setPodLabels([ ...podLabels$(), {
			key: undefined,
			value: undefined
		} ])
	}

	function addDeploymentLabel() {
		setDeploymentLabels([ ...deploymentLabels$(), {
			key: undefined,
			value: undefined
		} ])
	}

	function addNatchLabel() {
		setMatchLabels([ ...matchLabels$(), {
			key: undefined,
			value: undefined
		} ])
	}

	function render() {
		return <>
			<TextField
				required
				label="Deployment name"
				onChange={(e) => deploymentName = undefinedIfBlank(e.target.value)}
			/>

			{chipIfNotZero(deploymentLabels$().length, "Deployment labels")}
			{deploymentLabels$().map(c => <Label obj={c} />)}
			<Button onClick={addDeploymentLabel} variant="contained">Add deployment label</Button>

			{chipIfNotZero(matchLabels$().length, "Match labels")}
			{matchLabels$().map(c => <Label obj={c} />)}
			<Button onClick={addNatchLabel} variant="contained">Add match label</Button>

			<TextField
				required
				label="Replicas"
				onChange={(e) => {
					let val = undefinedIfBlank(e.target.value)
					if (val == undefined) {
						replicas = 0
					} else {
						replicas = +val
					}
				}}
			/>

			<InputLabel id="fadfasdf11">Deployment Strategy</InputLabel>
			<Select labelId="fadfasdf11" onChange={e => strategy = e.target.value as DeploymentStrategy}>
				<MenuItem value={DeploymentStrategy.Recreate}>Recreate</MenuItem>
				<MenuItem value={DeploymentStrategy.RollingUpdate}>RollingUpdate</MenuItem>
			</Select>

			<TextField
				required
				label="Namespace"
				defaultValue="default"
				onChange={(e) => namespace = undefinedIfBlank(e.target.value)}
			/>

			<TextField
				required
				label="Pod name"
				onChange={(e) => podName = undefinedIfBlank(e.target.value)}
			/>

			{chipIfNotZero(podLabels$().length, "Pod labels")}
			{podLabels$().map(c => <Label obj={c} />)}
			<Button onClick={addPodLabel} variant="contained">Add pod label</Button>

			{containers$().map(c => <Container obj={c} />)}
			<Button onClick={addContainer} variant="contained">Add container</Button>

			<Button onClick={create} variant="outlined">Create</Button>

			{renderResult(result$())}
		</>
	}

	return <>{render()}</>
}

function Delete(props: { type: 'pod' | 'deployment' }) {
	const { type } = props

	const [ result$, setResult ] = createSignal<Result>()

	let namespace: string | undefined = 'default'
	let name: string | undefined = undefined

	function execute() {
		if (namespace == undefined) {
			setResult({ success: false, error: "namespace is empty" })
			return
		}

		if (name == undefined) {
			setResult({ success: false, error: "name is empty" })
			return
		}

		const request: DeleteRequest = { name, namespace };

		void makeRequest('DELETE', `/v1/${type}`, request, setResult)
	}

	return <>
		<TextField
			required
			label="Namespace"
			defaultValue="default"
			onChange={(e) => namespace = undefinedIfBlank(e.target.value)}
		/>

		<TextField
			required
			label="Name"
			onChange={(e) => name = undefinedIfBlank(e.target.value)}
		/>

		<Button onClick={execute} variant="outlined">Delete</Button>

		{renderResult(result$())}
	</>
}

function App() {
	const [ type$, setType ] = createSignal<Type | undefined>(undefined);

	function changeType(newType: SelectChangeEvent) {
		setType(newType.target.value as Type)
	}

	return (
		<Box justifyContent="center" alignItems="center">
			<Stack sx={{ width: '100%', maxWidth: "500px" }} spacing={2}>
				<Select onChange={changeType} defaultOpen={true}>
					<MenuItem value={Type.Deployment}>Deployment</MenuItem>
					<MenuItem value={Type.DeleteDeployment}>Delete Deployment</MenuItem>
				</Select>

				{(() => {
					let type = type$()
					if (type == undefined) {
						return <></>
					} else {
						switch (type) {
							case Type.Deployment:
								return <Deployment />
							case Type.DeleteDeployment:
								return <Delete type="deployment" />
						}
					}
				})()}
			</Stack>
		</Box>
	)
}

export default App
