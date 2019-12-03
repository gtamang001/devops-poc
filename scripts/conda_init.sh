# added by Anaconda3 5.3.0 installer
# >>> conda init >>>
# !! Contents within this block are managed by 'conda init' !!
# run the conda shell.bash setup hook. It emits on stdout something that looks like shell code, save that code in __conda_setup.
__conda_setup="$(CONDA_REPORT_ERRORS=false '/home/pratyush/anaconda3/bin/conda' shell.bash hook 2> /dev/null)"
# if the setup step succeeded, execute the text in $__conda_setup as code. The \ means to ignore any defined aliases.
if [ $? -eq 0 ]; then 
    \eval "$__conda_setup"
# branch to run if `conda shell.bash hook` failed.
else 
    # if a conda configuration file exists, run it and tell conda to not change your PS1 shell display variable.
    if [ -f "/home/pratyush/anaconda3/etc/profile.d/conda.sh" ]; then
        . "/home/pratyush/anaconda3/etc/profile.d/conda.sh" 
        CONDA_CHANGEPS1=false conda activate base
    else
        # otherwise just change your PATH.
        \export PATH="/home/pratyush/anaconda3/bin:$PATH"
    fi
fi