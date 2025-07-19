basepath=$(cd `dirname $0`; pwd)
mkdir -p ${basepath}/moments
docker run --name moments -e JWT_KEY=cfqYVP6CZm9mSqLVGlmL -e PORT=37892 -d -v ${basepath}/moments:/app/data -p 37892:37892 kingwrcy/moments:latest