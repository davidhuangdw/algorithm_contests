package atcoder.a6

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toLong)

    val (N, K) = inputInts()
    val rev = K.toString().reversed().toLong()
    if(K > N || rev < K){
        println(0)
    }else{
        var cnt = 0
        for(v in hashSetOf(K, rev)){
            var cur = v
            while(cur <= N){
                cnt += 1
                cur *= 10
            }
        }
        println(cnt)
    }

}
