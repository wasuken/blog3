import * as React from 'react';
import * as ReactDOM from 'react-dom';

interface AppProps{
	
}
interface AppState{
	
}

class App extends React.Component<AppProps, AppState>{
	constructor(props: AppProps){
		super(props);
	}
	render(){
		return (
			<div>hello</div>
		)
	}
}

ReactDOM.render(
	<App />,
	document.querySelector('#app'),
);
