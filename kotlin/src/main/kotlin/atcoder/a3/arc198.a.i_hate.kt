package atcoder.a3
fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (N, ) = input_ints()
    if(N == 1){
        println("1\n1\n")
    }else{
        println(N/2);
        println((2..N step 2).toList().joinToString(" "))
    }
}
