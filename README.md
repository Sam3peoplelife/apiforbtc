# apiforbtc
API FOR BTC/UAH CURRENCY

Makefile - file for building
In store.go path for gmail.txt is full because host is running on the other directory

in apiserver sending emails is written but its not working cuz i have problem with pointer with getting Emails from file

DockerFile needs to look like this
# Use an existing docker image as a base
FROM debian:latest

# Set the working directory in the container
WORKDIR \apischool\cmd\apiserver

# Copy the current directory contents into the container at /app
COPY . /apiserver

# Install make
RUN apt-get update && apt-get install -y Makefile

# Run the make command
RUN Makefile

# application's run command
CMD ["./apiserver"]

docker build -t image

docker run -p 4000:80 image

Docker engine error that i cant resolve

Problem is in my Windows arm64, that doesnt support Docker engine. I tried to open Docker with VM, but got the same error.

Newtonsoft.Json.JsonReaderException:
Unexpected character encountered while parsing value: n. Path '', line 0, position 0.
   at Newtonsoft.Json.JsonTextReader.ParseValue()
   at Newtonsoft.Json.JsonReader.ReadForType(JsonContract contract, Boolean hasConverter)
   at Newtonsoft.Json.Serialization.JsonSerializerInternalReader.Deserialize(JsonReader reader, Type objectType, Boolean checkAdditionalContent)
   at Newtonsoft.Json.JsonSerializer.DeserializeInternal(JsonReader reader, Type objectType)
   at Newtonsoft.Json.JsonConvert.DeserializeObject(String value, Type type, JsonSerializerSettings settings)
   at Newtonsoft.Json.JsonConvert.DeserializeObject[T](String value, JsonSerializerSettings settings)
   at Docker.Engines.WSL2.LinuxWSL2Engine.<DoStartAsync>d__10.MoveNext() in C:\workspaces\PR-21718\src\github.com\docker\pinata\win\src\Docker.Engines\WSL2\LinuxWSL2Engine.cs:line 58
--- End of stack trace from previous location where exception was thrown ---
   at System.Runtime.ExceptionServices.ExceptionDispatchInfo.Throw()
   at System.Runtime.CompilerServices.TaskAwaiter.HandleNonSuccessAndDebuggerNotification(Task task)
   at Docker.ApiServices.StateMachines.TaskExtensions.<WrapAsyncInCancellationException>d__0.MoveNext() in C:\workspaces\PR-21718\src\github.com\docker\pinata\win\src\Docker.ApiServices\StateMachines\TaskExtensions.cs:line 29
--- End of stack trace from previous location where exception was thrown ---
   at System.Runtime.ExceptionServices.ExceptionDispatchInfo.Throw()
   at System.Runtime.CompilerServices.TaskAwaiter.HandleNonSuccessAndDebuggerNotification(Task task)
   at Docker.ApiServices.StateMachines.StartTransition.<DoRunAsync>d__7.MoveNext() in C:\workspaces\PR-21718\src\github.com\docker\pinata\win\src\Docker.ApiServices\StateMachines\StartTransition.cs:line 79
--- End of stack trace from previous location where exception was thrown ---
   at System.Runtime.ExceptionServices.ExceptionDispatchInfo.Throw()
   at Docker.ApiServices.StateMachines.StartTransition.<DoRunAsync>d__7.MoveNext() in C:\workspaces\PR-21718\src\github.com\docker\pinata\win\src\Docker.ApiServices\StateMachines\StartTransition.cs:line 108
--- End of stack trace from previous location where exception was thrown ---
   at System.Runtime.ExceptionServices.ExceptionDispatchInfo.Throw()
   at System.Runtime.CompilerServices.TaskAwaiter.HandleNonSuccessAndDebuggerNotification(Task task)
   at Docker.ApiServices.StateMachines.EngineStateMachine.<StartAsync>d__17.MoveNext() in C:\workspaces\PR-21718\src\github.com\docker\pinata\win\src\Docker.ApiServices\StateMachines\EngineStateMachine.cs:line 97
--- End of stack trace from previous location where exception was thrown ---
   at System.Runtime.ExceptionServices.ExceptionDispatchInfo.Throw()
   at System.Runtime.CompilerServices.TaskAwaiter.HandleNonSuccessAndDebuggerNotification(Task task)
   at Docker.Engines.Engines.<StartAsync>d__25.MoveNext() in C:\workspaces\PR-21718\src\github.com\docker\pinata\win\src\Docker.Engines\Engines.cs:line 126
