package luogu_template.string

fun manacher(ss: String): Int{
    var (j, jr) = -1 to -1

    val n = ss.length*2+3
    fun get(k: Int): Char{
        if(k == 0) return '@'
        if(k == n-1) return '$'
        return if(k%2==1) '*' else ss[k/2-1]
    }
//    val s = "@*" + ss.toList().joinToString("*") + "*$"
    val dis = IntArray(n)
    for(i in 1..n-2){
        var r = i
        if(jr >= i){
            r = minOf(jr, i+dis[2*j-i])
        }
        while(get(r+1) == get(2*i-(r+1))){
            r++
        }
        dis[i] = r - i
        if(r >= jr){
            jr = r
            j = i
        }
    }
    return dis.max()
}

fun main() {
    val s = readLine()!!
    println(manacher(s))
}
