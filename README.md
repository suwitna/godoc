# godoc
1. Install docker toolbox on Window 10 home
2. Clone data from github   ->git clone --recursive https://github.com/suwitna/godoc.git
3. Deploy container from project to our container ->docker build -t godoc .
4. Run our container ->docker run -d --publish 8089:8088 godoc
