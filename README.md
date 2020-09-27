# Project for manage and run acceptance tests

## Why?
* This project can help you to test your REST services using only one file with tests.
* Backend of this project run all your tests in parallel, so it's very quick
* You don't need to use some script language (such as JS or Python). Support and maintain its infrastructure can be very boring
* These tests don't require to be written by your software engineer - it can be QA or even (really!?) Manager
* It's open-source project (if you want to use it [locally](./deployment.md)) and free to use

## Getting started
* Sign up
* Create objects and commands
* Write the file with tests
* Upload it and run tests
* Analyze tests report

## Example of usage
### Example of service implementation (node.js)
```js
users.get('/:hash', (req, res) => {
  const user = usersRepository.find(u => u.hash === req.params.hash);

  if (user) {
    res.status(200).send({
      status: 'ok',
      data: user
    });
  } else {
    res.status(200).send({
      status: 'error',
      data: 'user-not-found'
    });
  }
});

users.post('/', (req, res) => {
  const { name } = req.body;
  const hash = uuid4();

  usersRepository.push({
    hash, name
  });
  res.status(200).send({status: 'ok', hash});
});
```
### Tests example
```
BEGIN
    createUserResponse = CREATE USER {"name": "Joe"}
    
    ASSERT createUserResponse.status EQUALS ok
    
    userResponse = GET USER ${createUserResponse.hash}
    
    ASSERT userResponse.status EQUALS ok
    ASSERT userResponse.data.name EQUALS Joe
END

BEGIN
    userResponse = GET USER not-exists-hash
    
    ASSERT userResponse.status EQUALS error
    ASSERT userResponse.data EQUALS user-not-found
END
```
