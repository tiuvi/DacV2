package shell

import (
	"errors"
	"os/exec"
	"os/user"
	"context"
	"time"
)

//En caso de que la aplicacion se lance sin permiso de administrador -> EXIT
func IsUserRoot()(isRoot bool , err error){

	user , err := user.Current()
	if err != nil{
		return
	}

	if(user.Name != "root"){
		return false , nil
	}

	return true , nil
}

//Simplifica las peticiones al cmd
func cmdOperation(name string,timeOut time.Duration,  arg ...string)(out []byte, err error){

	var cmd *exec.Cmd
	if timeOut != 0 {

		ctx, cancel := context.WithTimeout(context.Background(), timeOut)
		defer cancel()
		cmd = exec.CommandContext(ctx , name, arg... )

	}else{

		cmd = exec.Command( name, arg... )
	}

	out , err = cmd.Output()
	if err != nil {
		
		switch e := err.(type) {
	
			case *exec.Error:
				
				//println("Debug inError exeCmd " + name + " -> failed executing -> ", err)
				return out, err
			case *exec.ExitError:
				
				//println("Debug inError exeCmd " + name + " -> command exit rc = " ,e.ExitCode() , string(e.Stderr) )
				return out, errors.New(string(e.Stderr))
			default:
				//println("Debug inError exeCmd " + name + " -> SinInformacion -> " , err)
				return out, err
		}
	}

	return out, nil
}

//Simplifica las peticiones al cmd
func Cmd(name string, arg ...string)(out []byte, err error){

	return cmdOperation(name ,0,  arg... )

}

func CmdWithTimeOut(name string,timeOut time.Duration, arg ...string)(out []byte, err error){

	return cmdOperation(name ,timeOut,  arg... )

}

func CmdString(name string, arg ...string)( string,  error){

	out , err := Cmd(name , arg...)

	return string(out), err
}

func CmdStringWithTimeOut(name string,timeOut time.Duration, arg ...string)( string,  error){

	out , err := Cmd(name , arg...)

	return string(out), err
}



