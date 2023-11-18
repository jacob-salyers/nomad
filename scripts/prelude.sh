
if [ "$CODE_DIR" = '' ]
then 
	echo 'please export CODE_DIR' >&2
	exit 1
fi
if [ "$NOMAD_ROOT" = '' ]
then 
	echo 'please export NOMAD_ROOT' >&2
	exit 1
fi

dir="$CODE_DIR/nomad"

