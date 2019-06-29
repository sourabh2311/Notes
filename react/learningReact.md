- We need to use class based component in case we plan to use state
- If the state is changed, react will automatically update all the child components which are receiving that state values.

* list of possible events: https://reactjs.org/docs/events.html#mouse-events
* Whenever you want to change the state, use setState
* setState can as well be simple passed a new state, instead of passing it a function
* ```
    componentDidMount() {  // called when the component first gets mounted
        // GET the data I need to correctly display
    }

    shouldComponentUpdate(nextProps, nextState) {
        // return true if want it to update
        // return false if not
    }

    componentWillUnmount() {
        // teardown or cleanup your code before your component disappears
        // (E.g. remove event listeners)
    }
  ```

* For conditional rendering, you can use `&&` as well, as, `first && second` returns simply first if the first is false o/w it simply returns second
