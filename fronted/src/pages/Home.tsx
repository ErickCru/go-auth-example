
const Home = (props: { name: string }) => {

    return (
        <div>
            {props.name ? 'Bienvenido ' + props.name : 'You are not logged in'}
        </div>
    )
}

export default Home
