# If you wanna using other randomx fork, change the branch
# branches avaliable: master(=random-x) random-xl random-wow random-arq
echo "Building randomx..."

cd RandomX
mkdir build
cd build
cmake -G "Unix Makefiles" ..
make -j`nproc`
mv librandomx.a ../../lib
cd ..
rm -rf build
cd ..
